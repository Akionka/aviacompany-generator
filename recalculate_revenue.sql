DELIMITER $$

CREATE EVENT recalculate_revenue
ON SCHEDULE EVERY 1 DAY
STARTS TIMESTAMP(CURRENT_DATE, '00:00:00')
DO
BEGIN
    DECLARE exit handler for sqlexception
    BEGIN
        ROLLBACK;
        RESIGNAL;
    END;

    START TRANSACTION;

    INSERT INTO booking_office_revenue (booking_office_id, date, sum)
    SELECT * FROM (WITH RECURSIVE dates (date) AS
    (
        SELECT MIN(DATE(date)) FROM purchase
        UNION ALL
        SELECT date + INTERVAL 1 DAY FROM dates
        WHERE date + INTERVAL 1 DAY <= CURDATE()
    )
    SELECT /*+ SET_VAR(cte_max_recursion_depth = 1M) */
    p.booking_office_id as bio, d.date as date, sum(p.total_price) as tsum FROM dates d
    JOIN purchase p ON d.date = date(p.date)
    GROUP BY p.booking_office_id, d.date) t
    ON DUPLICATE KEY UPDATE sum = t.tsum;

    COMMIT;
END$$

DELIMITER ;