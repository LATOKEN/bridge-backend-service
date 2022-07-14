package storage

var (
	createTxTypeIfNotExists = `
        DO $$
        BEGIN
            IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tx_types') THEN
                CREATE TYPE tx_types AS ENUM
            ('PASSED', 'VOTE', 'SPEND', 'EXPIRED', 'UPDATE', 'DEPOSIT');
            END IF;
        END$$;
    `

	createTxLogStatusIfNotExists = `
        DO $$
        BEGIN
            IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tx_log_statuses') THEN
                CREATE TYPE tx_log_statuses AS ENUM
            ('INIT', 'CONFIRMED');
            END IF;
        END$$;
    `

	createTxStatusIfNotExists = `
        DO $$
        BEGIN
            IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tx_statuses') THEN
                CREATE TYPE tx_statuses AS ENUM
            ('INIT', 'NOT_FOUND', 'PENDING', 'FAILED', 'SUCCESS', 'LOST');
            END IF;
        END$$;
    `

	// sql script for 'block_type'
	createTypeIfNotExistsBlockTypeStatus = `
        DO $$
        BEGIN
            IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'block_type') THEN
                CREATE TYPE block_type AS ENUM
            ('CURRENT', 'PREVIOUS');
            END IF;
        END$$;
    `
)
