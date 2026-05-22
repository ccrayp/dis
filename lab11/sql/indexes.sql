-- Расширение для нечеткого поиска по строковым полям
CREATE EXTENSION IF NOT EXISTS pg_trgm;

--1. Индексы на внешние ключи (ускоряют JOIN и каскадные операции)

CREATE INDEX idx_users_host_id ON users(host_id);
CREATE INDEX idx_policies_user_id ON policies(user_id);
CREATE INDEX idx_requests_user_id ON requests(user_id);
CREATE INDEX idx_decisions_request_id ON decisions(request_id);
CREATE INDEX idx_decisions_policy_id ON decisions(policy_id);
CREATE INDEX idx_policy_contents_rule_id ON policy_contents(rule_id);
CREATE INDEX idx_policy_contents_service_id ON policy_contents(service_id);

-- 2. Индексы по времени (ускоряют выборки за определенный период)

CREATE INDEX idx_requests_timestamp ON requests(timestamp);
CREATE INDEX idx_decisions_timestamp ON decisions(timestamp);

-- 3. Индекс по статусу политик (ускоряет поиск активных/неактивных политик)

CREATE INDEX idx_policies_status ON policies(status);

-- 4. Индексы для нечеткого поиска (поиск по части слова и similarity())

CREATE INDEX idx_rules_rule_name_trgm ON rules USING GIN (rule_name gin_trgm_ops);
CREATE INDEX idx_users_user_name_trgm ON users USING GIN (user_name gin_trgm_ops);
CREATE INDEX idx_services_service_name_trgm ON services USING GIN (service_name gin_trgm_ops);
CREATE INDEX idx_policies_policy_name_trgm ON policies USING GIN (policy_name gin_trgm_ops);

-- 5. Индексы для аутентификации

CREATE INDEX idx_refresh_tokens_username ON refresh_tokens(token);
CREATE UNIQUE INDEX idx_sys_users_username ON sys_users(username);