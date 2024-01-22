package seeding

import (
	"fmt"
	"perftesting/db"
)

type SeederPartitionHashIndex struct {
	NumberOfPartitions int
}

func GetSeederPartitionHashIndex(tableName string, noOfPartitions int, columnarEngine bool) *Seeder {
	temp := &Seeder{}
	temp.TableName = tableName
	temp.BatchSize = 100000
	temp.ColumnEngine = columnarEngine
	temp.Schemer = &SeederPartitionHashIndex{
		NumberOfPartitions: noOfPartitions,
	}
	return temp
}

func (conf *SeederPartitionHashIndex) CreateSchema(s *Seeder) error {
	_, err := db.DB.Exec(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			validatorindex BIGINT,
			epoch BIGINT,
			attestations_source_reward BIGINT,
			attestations_target_reward BIGINT,
			attestations_head_reward BIGINT,
			attestations_inactivity_reward BIGINT,
			attestations_inclusion_reward BIGINT,
			attestations_reward BIGINT,
			attestations_ideal_source_reward BIGINT,
			attestations_ideal_target_reward BIGINT,
			attestations_ideal_head_reward BIGINT,
			attestations_ideal_inactivity_reward BIGINT,
			attestations_ideal_inclusion_reward BIGINT,
			attestations_ideal_reward BIGINT,
			blocks_scheduled INTEGER,
			blocks_proposed INTEGER,
			blocks_cl_reward BIGINT,
			blocks_el_reward NUMERIC,
			sync_scheduled INTEGER,
			sync_executed INTEGER,
			sync_rewards BIGINT,
			slashed BOOLEAN,
			balance_start BIGINT,
			balance_end BIGINT,
			deposits_count BIGINT,
			deposits_amount BIGINT,
			withdrawals_count BIGINT,
			withdrawals_amount BIGINT
		) PARTITION BY hash(validatorindex)
	`, s.TableName))
	if err != nil {
		return err
	}

	_, err = db.DB.Exec(fmt.Sprintf(`
		CREATE INDEX IF NOT EXISTS %s_validatorindex ON %[1]s (validatorindex, epoch)
	`, s.TableName))
	if err != nil {
		return err
	}

	for i := 0; i < conf.NumberOfPartitions; i++ {
		partitionCreate := fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s_%d PARTITION OF %[1]s
				FOR VALUES WITH (MODULUS %[3]d, REMAINDER %[2]d)
		`, s.TableName, i, conf.NumberOfPartitions)

		_, err = db.DB.Exec(partitionCreate)
		if err != nil {
			return err
		}

		// Column engine leaves
		if s.ColumnEngine {
			fmt.Printf("adding column engine to %[1]s_%[2]d\n", s.TableName, i)
			_, err = db.DB.Exec(fmt.Sprintf(`
				SELECT google_columnar_engine_add(
					relation => '%[1]s_%[2]d',
					columns => 'attestations_head_reward,attestations_source_reward,attestations_target_reward,blocks_cl_reward,epoch,validatorindex'
				);
				`, s.TableName, i))
			if err != nil {
				return err
			}
		}

	}

	return nil
}
