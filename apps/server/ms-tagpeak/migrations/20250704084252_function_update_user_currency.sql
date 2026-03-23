-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';


/**
 * Function: convert_user_transactions
 *
 * Description:
 *   This function retrieves all transactions for a specific user, converts the
 *   amount_user and commission_user values to the specified target currency,
 *   and updates the transaction table with the converted values.
 *   If a transaction has a related reward, it also updates the current_reward_user
 *   value in the reward table.
 *   The conversion is done using the formula from GetAmountByCurrencyRate: (value * rate) / rateSource
 *   Note: Currency exchange rates are always coming from Euro, so we use amount_target and commission_target
 *   for the calculation instead of amount_source and commission_source.
 *
 * Parameters:
 *   user_id - TEXT of the user whose transactions to retrieve and update
 *   target_currency_code - Currency code to convert the amounts to (e.g., 'USD', 'EUR')
 *
 * Returns:
 *   A table with the following columns:
 *     - transaction_uuid: UUID of the transaction
 *     - original_amount_user: Original amount in user's currency
 *     - converted_amount_user: Amount converted to target currency
 *     - original_commission_user: Original commission in user's currency
 *     - converted_commission_user: Commission converted to target currency
 */
CREATE OR REPLACE FUNCTION convert_user_transactions(user_id TEXT, target_currency_code TEXT)
	RETURNS TABLE (
		transaction_uuid UUID,
		original_amount_user FLOAT,
		converted_amount_user FLOAT,
		original_commission_user FLOAT,
		converted_commission_user FLOAT
	) AS $$
DECLARE
	transaction_count INTEGER;
	reward_count INTEGER;
	transaction_record RECORD;
	reward_record RECORD;
	converted_amount FLOAT;
	converted_commission FLOAT;
	converted_reward FLOAT;
BEGIN
	-- Log function start
	RAISE NOTICE 'Starting convert_user_transactions for user_id: % to currency: %', user_id, target_currency_code;

	-- Count transactions for logging
	SELECT COUNT(*) INTO transaction_count
	FROM transaction t
	WHERE t."user" = user_id;

	RAISE NOTICE 'Found % transactions for user', transaction_count;

	-- Create a temporary table to store transactions for logging and processing
	CREATE TEMP TABLE temp_transactions AS
	SELECT
		t.uuid,
		t.amount_user,
		t.amount_target,
		t.commission_user,
		t.commission_target,
		t.currency_source,
		t.currency_target,
		cer.rates,
		cer.base
	FROM
		transaction t
			JOIN
		currency_exchange_rate cer ON t.currency_exchange_rate_uuid = cer.uuid
	WHERE
		t."user" = user_id;

	-- Log each transaction's initial values
	FOR transaction_record IN SELECT * FROM temp_transactions LOOP
			RAISE NOTICE 'Transaction %: Initial amount_user=%, commission_user=%, currency=%',
				transaction_record.uuid,
				transaction_record.amount_user,
				transaction_record.commission_user,
				transaction_record.currency_source;
		END LOOP;

	-- Count rewards for logging
	SELECT COUNT(*) INTO reward_count
	FROM reward r
		     JOIN transaction t ON r.transaction_uuid = t.uuid
	WHERE t."user" = user_id;

	RAISE NOTICE 'Found % rewards associated with user transactions', reward_count;

	-- Process each transaction and update the transaction table
	FOR transaction_record IN SELECT * FROM temp_transactions LOOP
			-- Calculate converted amount
			IF transaction_record.base = target_currency_code THEN
				converted_amount := (transaction_record.amount_target * 1.0) / (transaction_record.rates->>transaction_record.currency_target)::FLOAT;
			ELSIF transaction_record.currency_target = target_currency_code THEN
				converted_amount := transaction_record.amount_target;
			ELSE
				converted_amount := (transaction_record.amount_target * (transaction_record.rates->>target_currency_code)::FLOAT) / (transaction_record.rates->>transaction_record.currency_target)::FLOAT;
			END IF;

			-- Calculate converted commission
			IF transaction_record.base = target_currency_code THEN
				converted_commission := (transaction_record.commission_target * 1.0) / (transaction_record.rates->>transaction_record.currency_target)::FLOAT;
			ELSIF transaction_record.currency_target = target_currency_code THEN
				converted_commission := transaction_record.commission_target;
			ELSE
				converted_commission := (transaction_record.commission_target * (transaction_record.rates->>target_currency_code)::FLOAT) / (transaction_record.rates->>transaction_record.currency_target)::FLOAT;
			END IF;

			-- Update the transaction table
			UPDATE transaction
			SET
				amount_user = converted_amount,
				commission_user = converted_commission,
				updated_at = NOW()
			WHERE uuid = transaction_record.uuid;

			-- Log the transaction conversion
			RAISE NOTICE 'Transaction %: Updated amount_user=% → %, commission_user=% → %',
				transaction_record.uuid,
				transaction_record.amount_user,
				converted_amount,
				transaction_record.commission_user,
				converted_commission;

			-- Check if this transaction has a related reward
			FOR reward_record IN
				SELECT
					r.uuid,
					r.current_reward_user,
					r.current_reward_target
				FROM
					reward r
				WHERE
					r.transaction_uuid = transaction_record.uuid
				LOOP
					-- Calculate converted reward
					IF transaction_record.base = target_currency_code THEN
						converted_reward := (reward_record.current_reward_target * 1.0) / (transaction_record.rates->>transaction_record.currency_target)::FLOAT;
					ELSIF transaction_record.currency_target = target_currency_code THEN
						converted_reward := reward_record.current_reward_target;
					ELSE
						converted_reward := (reward_record.current_reward_target * (transaction_record.rates->>target_currency_code)::FLOAT) / (transaction_record.rates->>transaction_record.currency_target)::FLOAT;
					END IF;

					-- Update the reward table
					UPDATE reward
					SET
						current_reward_user = converted_reward,
						currency_user = target_currency_code,
						updated_at = NOW()
					WHERE uuid = reward_record.uuid;

					-- Log the reward conversion
					RAISE NOTICE 'Reward % (for Transaction %): Updated current_reward_user=% → %',
						reward_record.uuid,
						transaction_record.uuid,
						reward_record.current_reward_user,
						converted_reward;
				END LOOP;
		END LOOP;

	-- Return the converted transactions
	RETURN QUERY
		SELECT
			ut.uuid,
			ut.amount_user AS original_amount_user,
			CASE
				WHEN ut.base = target_currency_code THEN
					(ut.amount_target * 1.0) / (ut.rates->>ut.currency_target)::FLOAT
				WHEN ut.currency_target = target_currency_code THEN
					ut.amount_target
				ELSE
					(ut.amount_target * (ut.rates->>target_currency_code)::FLOAT) / (ut.rates->>ut.currency_target)::FLOAT
				END AS converted_amount_user,
			ut.commission_user AS original_commission_user,
			CASE
				WHEN ut.base = target_currency_code THEN
					(ut.commission_target * 1.0) / (ut.rates->>ut.currency_target)::FLOAT
				WHEN ut.currency_target = target_currency_code THEN
					ut.commission_target
				ELSE
					(ut.commission_target * (ut.rates->>target_currency_code)::FLOAT) / (ut.rates->>ut.currency_target)::FLOAT
				END AS converted_commission_user
		FROM
			temp_transactions ut;

	-- Clean up temporary table
	DROP TABLE temp_transactions;

	-- Log function end
	RAISE NOTICE 'Finished convert_user_transactions for user_id: % to currency: %', user_id, target_currency_code;
END;
$$ LANGUAGE plpgsql;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

-- Restore the previous version of the function
CREATE OR REPLACE FUNCTION convert_user_transactions(user_id TEXT, target_currency_code TEXT)
	RETURNS TABLE (
		transaction_uuid UUID,
		original_amount_user FLOAT,
		converted_amount_user FLOAT,
		original_commission_user FLOAT,
		converted_commission_user FLOAT
	) AS $$
DECLARE
	transaction_count INTEGER;
	reward_count INTEGER;
	transaction_record RECORD;
	reward_record RECORD;
	converted_amount FLOAT;
	converted_commission FLOAT;
	converted_reward FLOAT;
BEGIN
	-- Log function start
	RAISE NOTICE 'Starting convert_user_transactions for user_id: % to currency: %', user_id, target_currency_code;

	-- Count transactions for logging
	SELECT COUNT(*) INTO transaction_count
	FROM transaction t
	WHERE t."user" = user_id;

	RAISE NOTICE 'Found % transactions for user', transaction_count;

	-- Create a temporary table to store transactions for logging and processing
	CREATE TEMP TABLE temp_transactions AS
	SELECT
		t.uuid,
		t.amount_user,
		t.amount_source,
		t.commission_user,
		t.commission_source,
		t.currency_source,
		cer.rates,
		cer.base
	FROM
		transaction t
			JOIN
		currency_exchange_rate cer ON t.currency_exchange_rate_uuid = cer.uuid
	WHERE
		t."user" = user_id;

	-- Log each transaction's initial values
	FOR transaction_record IN SELECT * FROM temp_transactions LOOP
			RAISE NOTICE 'Transaction %: Initial amount_user=%, commission_user=%, currency=%',
				transaction_record.uuid,
				transaction_record.amount_user,
				transaction_record.commission_user,
				transaction_record.currency_source;
		END LOOP;

	-- Count rewards for logging
	SELECT COUNT(*) INTO reward_count
	FROM reward r
		     JOIN transaction t ON r.transaction_uuid = t.uuid
	WHERE t."user" = user_id;

	RAISE NOTICE 'Found % rewards associated with user transactions', reward_count;

	-- Process each transaction and update the transaction table
	FOR transaction_record IN SELECT * FROM temp_transactions LOOP
			-- Calculate converted amount
			IF transaction_record.base = target_currency_code THEN
				converted_amount := (transaction_record.amount_source * 1.0) / (transaction_record.rates->>transaction_record.currency_source)::FLOAT;
			ELSIF transaction_record.currency_source = target_currency_code THEN
				converted_amount := transaction_record.amount_source;
			ELSE
				converted_amount := (transaction_record.amount_source * (transaction_record.rates->>target_currency_code)::FLOAT) / (transaction_record.rates->>transaction_record.currency_source)::FLOAT;
			END IF;

			-- Calculate converted commission
			IF transaction_record.base = target_currency_code THEN
				converted_commission := (transaction_record.commission_source * 1.0) / (transaction_record.rates->>transaction_record.currency_source)::FLOAT;
			ELSIF transaction_record.currency_source = target_currency_code THEN
				converted_commission := transaction_record.commission_source;
			ELSE
				converted_commission := (transaction_record.commission_source * (transaction_record.rates->>target_currency_code)::FLOAT) / (transaction_record.rates->>transaction_record.currency_source)::FLOAT;
			END IF;

			-- Update the transaction table
			UPDATE transaction
			SET
				amount_user = converted_amount,
				commission_user = converted_commission,
				updated_at = NOW()
			WHERE uuid = transaction_record.uuid;

			-- Log the transaction conversion
			RAISE NOTICE 'Transaction %: Updated amount_user=% → %, commission_user=% → %',
				transaction_record.uuid,
				transaction_record.amount_user,
				converted_amount,
				transaction_record.commission_user,
				converted_commission;

			-- Check if this transaction has a related reward
			FOR reward_record IN
				SELECT
					r.uuid,
					r.current_reward_user,
					r.current_reward_source
				FROM
					reward r
				WHERE
					r.transaction_uuid = transaction_record.uuid
				LOOP
					-- Calculate converted reward
					IF transaction_record.base = target_currency_code THEN
						converted_reward := (reward_record.current_reward_source * 1.0) / (transaction_record.rates->>transaction_record.currency_source)::FLOAT;
					ELSIF transaction_record.currency_source = target_currency_code THEN
						converted_reward := reward_record.current_reward_source;
					ELSE
						converted_reward := (reward_record.current_reward_source * (transaction_record.rates->>target_currency_code)::FLOAT) / (transaction_record.rates->>transaction_record.currency_source)::FLOAT;
					END IF;

					-- Update the reward table
					UPDATE reward
					SET
						current_reward_user = converted_reward,
						currency_user = target_currency_code,
						updated_at = NOW()
					WHERE uuid = reward_record.uuid;

					-- Log the reward conversion
					RAISE NOTICE 'Reward % (for Transaction %): Updated current_reward_user=% → %',
						reward_record.uuid,
						transaction_record.uuid,
						reward_record.current_reward_user,
						converted_reward;
				END LOOP;
		END LOOP;

	-- Return the converted transactions
	RETURN QUERY
		SELECT
			ut.uuid,
			ut.amount_user AS original_amount_user,
			CASE
				WHEN ut.base = target_currency_code THEN
					(ut.amount_source * 1.0) / (ut.rates->>ut.currency_source)::FLOAT
				WHEN ut.currency_source = target_currency_code THEN
					ut.amount_source
				ELSE
					(ut.amount_source * (ut.rates->>target_currency_code)::FLOAT) / (ut.rates->>ut.currency_source)::FLOAT
				END AS converted_amount_user,
			ut.commission_user AS original_commission_user,
			CASE
				WHEN ut.base = target_currency_code THEN
					(ut.commission_source * 1.0) / (ut.rates->>ut.currency_source)::FLOAT
				WHEN ut.currency_source = target_currency_code THEN
					ut.commission_source
				ELSE
					(ut.commission_source * (ut.rates->>target_currency_code)::FLOAT) / (ut.rates->>ut.currency_source)::FLOAT
				END AS converted_commission_user
		FROM
			temp_transactions ut;

	-- Clean up temporary table
	DROP TABLE temp_transactions;

	-- Log function end
	RAISE NOTICE 'Finished convert_user_transactions for user_id: % to currency: %', user_id, target_currency_code;
END;
$$ LANGUAGE plpgsql;

-- +goose StatementEnd
