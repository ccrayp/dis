-- Логи с пояснеиями
CREATE VIEW extended_logs AS
SELECT
	ROW_NUMBER() OVER() AS number,
	u.user_name,
	r.context,
	p.policy_name,
	rl.condition,
	d.result
FROM decisions AS d
JOIN requests AS r
	USING (request_id)
JOIN users AS u
	USING (user_id)
JOIN policies AS p
	USING(policy_id)
JOIN policy_contents AS pc
	USING(policy_id)
JOIN rules AS rl
	USING(rule_id)

-- Список пользователей по хостам
CREATE VIEW host_users AS
SELECT
	h.ip,
	ARRAY_AGG(u.user_name ORDER BY u.user_name)
FROM users AS u
JOIN hosts AS h
	USING(host_id)
GROUP BY h.host_id
ORDER BY h.ip

-- Количество правил и сервисов в каждой политике
CREATE OR REPLACE VIEW policy_info AS
SELECT
	p.policy_id,
	p.policy_name,
	COUNT(pc.rule_id) AS count_of_rules,
	COUNT(pc.service_id) AS count_of_services
FROM policy_contents AS pc
JOIN policies AS p
	USING(policy_id)
GROUP BY p.policy_id, p.policy_name
ORDER BY p.policy_id

-- Количество запросов с хоста за полгода
CREATE VIEW hosts_requests_info AS
SELECT 
	h.host_id,
    h.ip,
    COUNT(u.user_id) AS users_count,
    COUNT(r.request_id) AS requests_count
FROM hosts AS h
JOIN users AS u
    USING (host_id)
JOIN requests AS r
    USING (user_id)
JOIN decisions AS d
    USING (request_id)
WHERE r.timestamp BETWEEN NOW() - INTERVAL '6 month' AND NOW()
GROUP BY h.host_id, h.ip
ORDER BY h.ip;

-- Пользователи без политик
CREATE VIEW users_without_policies AS
SELECT
    u.user_id,
    u.user_name,
    h.ip
FROM users AS u
JOIN hosts AS h
    USING (host_id)
LEFT JOIN policies AS p
    USING (user_id)
WHERE p.policy_id IS NULL
ORDER BY u.user_name;

-- Статистика по пользователям
CREATE VIEW user_decision_stats AS
SELECT
    u.user_name,
    COUNT(*) AS total_requests,
    COUNT(*) FILTER (WHERE d.result = TRUE) AS allowed,
    COUNT(*) FILTER (WHERE d.result = FALSE) AS denied
FROM decisions AS d
JOIN requests AS r
    USING (request_id)
JOIN users AS u
    USING (user_id)
GROUP BY u.user_id, u.user_name
ORDER BY total_requests DESC;

-- Популятность сервисов
CREATE VIEW popular_services AS
SELECT
    s.service_name,
    COUNT(*) AS usage_count
FROM policy_contents AS pc
JOIN services AS s
    USING (service_id)
GROUP BY s.service_id, s.service_name
ORDER BY usage_count DESC, s.service_name;

-- Процент отказа пользователям
CREATE VIEW user_denial_rate AS
SELECT
    u.user_name,
    ROUND(
        100.0 * COUNT(*) FILTER (WHERE d.result = FALSE) / COUNT(*),
        2
    ) AS denial_percent
FROM decisions AS d
JOIN requests AS r
    USING (request_id)
JOIN users AS u
    USING (user_id)
GROUP BY u.user_id, u.user_name
ORDER BY denial_percent DESC;