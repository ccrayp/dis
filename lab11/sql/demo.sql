INSERT INTO hosts (ip) VALUES
('192.168.1.10'), ('192.168.1.11'), ('192.168.1.12'), ('192.168.1.13'), ('192.168.1.14'),
('192.168.1.15'), ('192.168.1.16'), ('192.168.1.17'), ('192.168.1.18'), ('192.168.1.19'),
('192.168.2.20'), ('192.168.2.21'), ('192.168.2.22'), ('192.168.2.23'), ('192.168.2.24'),
('10.0.0.25'), ('10.0.0.26'), ('10.0.0.27'), ('10.0.0.28'), ('10.0.0.29'),
('172.16.0.30'), ('172.16.0.31'), ('172.16.0.32'), ('172.16.0.33'), ('172.16.0.34');

INSERT INTO users (user_name, uid, gid, host_id) VALUES
('ivan.petrov', 1001, 1001, 1), ('maria.sidorova', 1002, 1002, 1), ('alexey.kozlov', 1003, 1003, 2),
('elena.novikova', 1004, 1004, 2), ('dmitry.volkov', 1005, 1005, 3), ('anna.morozova', 1006, 1006, 3),
('sergey.vasiliev', 1007, 1007, 4), ('olga.pavlova', 1008, 1008, 4), ('andrey.sokolov', 1009, 1009, 5),
('tatiana.mikhailova', 1010, 1010, 5), ('nikolai.fedorov', 1011, 1011, 6), ('ekaterina.egorova', 1012, 1012, 6),
('vladimir.kuznetsov', 1013, 1013, 7), ('natalia.popova', 1014, 1014, 7), ('pavel.aleksandrov', 1015, 1015, 8),
('irina.nikolaeva', 1016, 1016, 8), ('mikhail.andreev', 1017, 1017, 9), ('svetlana.makarova', 1018, 1018, 9),
('artem.romanov', 1019, 1019, 10), ('yulia.zakharova', 1020, 1020, 10), ('denis.ivanov', 1021, 1021, 11),
('elizaveta.smirnova', 1022, 1022, 11), ('maxim.kuzmin', 1023, 1023, 12), ('daria.lebedeva', 1024, 1024, 12),
('oleg.krylov', 1025, 1025, 13), ('valentina.orlova', 1026, 1026, 13), ('vladislav.ermakov', 1027, 1027, 14),
('alina.titova', 1028, 1028, 14), ('ruslan.mironov', 1029, 1029, 15), ('vera.belova', 1030, 1030, 15),
('gleb.sergeev', 1031, 1031, 16), ('nadezhda.komarova', 1032, 1032, 16), ('stanislav.medvedev', 1033, 1033, 17),
('ludmila.borisova', 1034, 1034, 17), ('igor.dmitriev', 1035, 1035, 18), ('alla.koroleva', 1036, 1036, 18),
('roman.belousov', 1037, 1037, 19), ('inga.terenteva', 1038, 1038, 19), ('stepan.gromov', 1039, 1039, 20),
('marina.melnikova', 1040, 1040, 20), ('vadim.nikonov', 1041, 1041, 21), ('vera.sorokina', 1042, 1042, 21),
('gennady.rodionov', 1043, 1043, 22), ('olga.ershova', 1044, 1044, 22), ('kirill.zinoviev', 1045, 1045, 23),
('lyubov.kalinina', 1046, 1046, 23), ('artur.bragin', 1047, 1047, 24), ('lilia.guseva', 1048, 1048, 24),
('vyacheslav.lazarev', 1049, 1049, 25), ('angelina.moiseeva', 1050, 1050, 25),
('roman.karpov', 1051, 1051, 1), ('diana.litvinova', 1052, 1052, 2), ('leonid.samsonov', 1053, 1053, 3),
('anastasia.zhukova', 1054, 1054, 4), ('yuri.bykov', 1055, 1055, 5), ('karina.danilova', 1056, 1056, 6),
('vitaly.konstantinov', 1057, 1057, 7), ('larisa.romanenko', 1058, 1058, 8), ('arthur.yakovlev', 1059, 1059, 9),
('elvira.frolova', 1060, 1060, 10), ('daniil.martynov', 1061, 1061, 11), ('zlata.kravtsova', 1062, 1062, 12),
('ilya.shcherbakov', 1063, 1063, 13), ('eva.fokina', 1064, 1064, 14), ('egor.baranov', 1065, 1065, 15),
('zoya.tyurina', 1066, 1066, 16), ('david.ishakov', 1067, 1067, 17), ('yana.galkina', 1068, 1068, 18),
('rodion.bespalov', 1069, 1069, 19), ('alla.shilova', 1070, 1070, 20), ('timur.subbotin', 1071, 1071, 21),
('oksana.mironova', 1072, 1072, 22), ('fedot.rybakov', 1073, 1073, 23), ('inga.panova', 1074, 1074, 24),
('svyatoslav.kozlov', 1075, 1075, 25), ('rada.petukhova', 1076, 1076, 1), ('lavr.ermolaev', 1077, 1077, 2),
('maya.tishchenko', 1078, 1078, 3), ('kuzma.kuzmin', 1079, 1079, 4), ('ada.samoylenko', 1080, 1080, 5);

INSERT INTO sys_users (username, password_hash) VALUES
('admin', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi'), -- password: password
('auditor', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi'),
('security', '$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi');

INSERT INTO services (service_name) VALUES
('ssh'), ('httpd'), ('nginx'), ('postgresql'), ('mysql'), ('redis'),
('docker'), ('kubernetes'), ('systemd'), ('cron'), ('rsyslog'), ('sshd'),
('ftp'), ('smtp'), ('dns'), ('dhcp'), ('nfs'), ('samba'), ('jenkins'),
('gitlab'), ('prometheus'), ('grafana'), ('elasticsearch'), ('kibana');

INSERT INTO rules (rule_name, condition, effect) VALUES
('Разрешить в пятницу днем', '[{"operator": "in", "type": "weekday", "value": ["friday"]}, {"operator": "between", "type": "timestamp", "value": "13:00-17:00"}]', true),
('Запретить в выходные', '[{"operator": "in", "type": "weekday", "value": ["saturday", "sunday"]}]', false),
('Только из офисной сети', '[{"operator": "in", "type": "ip", "value": "192.168.1.0/24"}]', true),
('Запретить из DMZ', '[{"operator": "in", "type": "ip", "value": "10.0.0.0/24"}]', false),
('Разрешить администраторам', '[{"operator": "contains", "type": "groups", "value": "sudo"}]', true),
('Только для пользователей с gid 1000-1050', '[{"operator": "between", "type": "gid", "value": "1000-1050"}]', true),
('Разрешить через VPN', '[{"operator": "regex", "type": "hostname", "value": "vpn-.*"}]', true),
('Запретить из гостевой сети', '[{"operator": "in", "type": "ip", "value": "172.16.0.0/24"}]', false),
('Разрешить утром', '[{"operator": "between", "type": "timestamp", "value": "06:00-12:00"}]', true),
('Запретить ночью', '[{"operator": "between", "type": "timestamp", "value": "23:00-05:00"}]', false);

INSERT INTO policies (policy_name, user_id, status) VALUES
('Default SSH Policy', 1, true), ('Admin Network Policy', 2, true), ('Restricted Access', 3, false),
('Developer Policy', 4, true), ('Night Shift Policy', 5, true), ('Weekend Policy', 6, false),
('VPN Only Policy', 7, true), ('Office Hours', 8, true), ('Production Access', 9, true),
('Testing Policy', 10, false);

INSERT INTO policy_contents (policy_id, service_id, rule_id) VALUES
(1, 1, 1), (1, 2, 3), (2, 3, 5), (2, 4, 6), (3, 1, 2), (3, 5, 8),
(4, 6, 4), (4, 7, 7), (5, 8, 9), (5, 9, 10), (6, 10, 2), (6, 11, 8),
(7, 12, 7), (7, 13, 3), (8, 14, 1), (8, 15, 9), (9, 16, 5), (9, 17, 6),
(10, 18, 4), (10, 19, 10);

INSERT INTO requests (user_id, context, timestamp) VALUES
(1, '{"host": {"ip": "192.168.1.10", "hostname": "workstation-1"}, "time": {"weekday": "monday", "timestamp": "2026-04-15T09:30:00Z"}, "user": {"gid": 1001, "uid": 1001, "groups": ["developers", "sudo"], "username": "ivan.petrov"}, "service": "ssh"}', '2026-04-15 09:30:00'),
(2, '{"host": {"ip": "192.168.1.11", "hostname": "workstation-2"}, "time": {"weekday": "tuesday", "timestamp": "2026-04-16T14:20:00Z"}, "user": {"gid": 1002, "uid": 1002, "groups": ["developers"], "username": "maria.sidorova"}, "service": "postgresql"}', '2026-04-16 14:20:00'),
(3, '{"host": {"ip": "10.0.0.25", "hostname": "dmz-server"}, "time": {"weekday": "wednesday", "timestamp": "2026-04-17T22:15:00Z"}, "user": {"gid": 1003, "uid": 1003, "groups": ["operators"], "username": "alexey.kozlov"}, "service": "nginx"}', '2026-04-17 22:15:00'),
(4, '{"host": {"ip": "192.168.2.20", "hostname": "dev-server"}, "time": {"weekday": "thursday", "timestamp": "2026-04-18T11:45:00Z"}, "user": {"gid": 1004, "uid": 1004, "groups": ["developers", "docker"], "username": "elena.novikova"}, "service": "docker"}', '2026-04-18 11:45:00'),
(5, '{"host": {"ip": "172.16.0.30", "hostname": "guest-wifi"}, "time": {"weekday": "friday", "timestamp": "2026-04-19T16:30:00Z"}, "user": {"gid": 1005, "uid": 1005, "groups": ["guests"], "username": "dmitry.volkov"}, "service": "ssh"}', '2026-04-19 16:30:00'),
(6, '{"host": {"ip": "192.168.1.12", "hostname": "workstation-3"}, "time": {"weekday": "saturday", "timestamp": "2026-04-20T03:00:00Z"}, "user": {"gid": 1006, "uid": 1006, "groups": ["developers"], "username": "anna.morozova"}, "service": "redis"}', '2026-04-20 03:00:00'),
(7, '{"host": {"ip": "192.168.1.13", "hostname": "admin-pc"}, "time": {"weekday": "sunday", "timestamp": "2026-04-21T08:00:00Z"}, "user": {"gid": 1007, "uid": 1007, "groups": ["admins", "sudo"], "username": "sergey.vasiliev"}, "service": "systemd"}', '2026-04-21 08:00:00'),
(8, '{"host": {"ip": "192.168.1.14", "hostname": "workstation-4"}, "time": {"weekday": "monday", "timestamp": "2026-04-22T13:30:00Z"}, "user": {"gid": 1008, "uid": 1008, "groups": ["developers"], "username": "olga.pavlova"}, "service": "mysql"}', '2026-04-22 13:30:00'),
(9, '{"host": {"ip": "10.0.0.26", "hostname": "database-1"}, "time": {"weekday": "tuesday", "timestamp": "2026-04-23T23:45:00Z"}, "user": {"gid": 1009, "uid": 1009, "groups": ["dba"], "username": "andrey.sokolov"}, "service": "postgresql"}', '2026-04-23 23:45:00'),
(10, '{"host": {"ip": "192.168.2.21", "hostname": "test-server"}, "time": {"weekday": "wednesday", "timestamp": "2026-04-24T10:15:00Z"}, "user": {"gid": 1010, "uid": 1010, "groups": ["testers"], "username": "tatiana.mikhailova"}, "service": "jenkins"}', '2026-04-24 10:15:00'),
(11, '{"host": {"ip": "172.16.0.31", "hostname": "vpn-gateway"}, "time": {"weekday": "thursday", "timestamp": "2026-04-25T07:30:00Z"}, "user": {"gid": 1011, "uid": 1011, "groups": ["remote"], "username": "nikolai.fedorov"}, "service": "ssh"}', '2026-04-25 07:30:00'),
(12, '{"host": {"ip": "192.168.1.15", "hostname": "build-server"}, "time": {"weekday": "friday", "timestamp": "2026-04-26T15:00:00Z"}, "user": {"gid": 1012, "uid": 1012, "groups": ["ci-cd", "docker"], "username": "ekaterina.egorova"}, "service": "docker"}', '2026-04-26 15:00:00'),
(13, '{"host": {"ip": "192.168.1.16", "hostname": "monitoring"}, "time": {"weekday": "saturday", "timestamp": "2026-04-27T12:00:00Z"}, "user": {"gid": 1013, "uid": 1013, "groups": ["monitoring"], "username": "vladimir.kuznetsov"}, "service": "prometheus"}', '2026-04-27 12:00:00'),
(14, '{"host": {"ip": "192.168.2.22", "hostname": "staging"}, "time": {"weekday": "sunday", "timestamp": "2026-04-28T18:30:00Z"}, "user": {"gid": 1014, "uid": 1014, "groups": ["developers"], "username": "natalia.popova"}, "service": "nginx"}', '2026-04-28 18:30:00'),
(15, '{"host": {"ip": "10.0.0.27", "hostname": "cache-1"}, "time": {"weekday": "monday", "timestamp": "2026-04-29T02:00:00Z"}, "user": {"gid": 1015, "uid": 1015, "groups": ["operators"], "username": "pavel.aleksandrov"}, "service": "redis"}', '2026-04-29 02:00:00'),
(16, '{"host": {"ip": "192.168.1.17", "hostname": "workstation-5"}, "time": {"weekday": "tuesday", "timestamp": "2026-04-30T09:00:00Z"}, "user": {"gid": 1016, "uid": 1016, "groups": ["designers"], "username": "irina.nikolaeva"}, "service": "httpd"}', '2026-04-30 09:00:00'),
(17, '{"host": {"ip": "172.16.0.32", "hostname": "public-wifi"}, "time": {"weekday": "wednesday", "timestamp": "2026-05-01T20:15:00Z"}, "user": {"gid": 1017, "uid": 1017, "groups": ["guests"], "username": "mikhail.andreev"}, "service": "ftp"}', '2026-05-01 20:15:00'),
(18, '{"host": {"ip": "192.168.2.23", "hostname": "backup-server"}, "time": {"weekday": "thursday", "timestamp": "2026-05-02T01:30:00Z"}, "user": {"gid": 1018, "uid": 1018, "groups": ["backup"], "username": "svetlana.makarova"}, "service": "rsyslog"}', '2026-05-02 01:30:00'),
(19, '{"host": {"ip": "192.168.1.18", "hostname": "workstation-6"}, "time": {"weekday": "friday", "timestamp": "2026-05-03T14:45:00Z"}, "user": {"gid": 1019, "uid": 1019, "groups": ["developers", "sudo"], "username": "artem.romanov"}, "service": "ssh"}', '2026-05-03 14:45:00'),
(20, '{"host": {"ip": "10.0.0.28", "hostname": "app-server"}, "time": {"weekday": "saturday", "timestamp": "2026-05-04T11:00:00Z"}, "user": {"gid": 1020, "uid": 1020, "groups": ["developers"], "username": "yulia.zakharova"}, "service": "gitlab"}', '2026-05-04 11:00:00'),
(21, '{"host": {"ip": "192.168.1.19", "hostname": "workstation-7"}, "time": {"weekday": "sunday", "timestamp": "2026-05-05T19:00:00Z"}, "user": {"gid": 1021, "uid": 1021, "groups": ["qa"], "username": "denis.ivanov"}, "service": "jenkins"}', '2026-05-05 19:00:00'),
(22, '{"host": {"ip": "172.16.0.33", "hostname": "vpn-usa"}, "time": {"weekday": "monday", "timestamp": "2026-05-06T06:00:00Z"}, "user": {"gid": 1022, "uid": 1022, "groups": ["remote"], "username": "elizaveta.smirnova"}, "service": "ssh"}', '2026-05-06 06:00:00'),
(23, '{"host": {"ip": "192.168.2.24", "hostname": "devops-server"}, "time": {"weekday": "tuesday", "timestamp": "2026-05-07T16:30:00Z"}, "user": {"gid": 1023, "uid": 1023, "groups": ["devops", "docker", "sudo"], "username": "maxim.kuzmin"}, "service": "kubernetes"}', '2026-05-07 16:30:00'),
(24, '{"host": {"ip": "192.168.1.20", "hostname": "workstation-8"}, "time": {"weekday": "wednesday", "timestamp": "2026-05-08T10:30:00Z"}, "user": {"gid": 1024, "uid": 1024, "groups": ["developers"], "username": "daria.lebedeva"}, "service": "postgresql"}', '2026-05-08 10:30:00'),
(25, '{"host": {"ip": "10.0.0.29", "hostname": "analytics"}, "time": {"weekday": "thursday", "timestamp": "2026-05-09T22:00:00Z"}, "user": {"gid": 1025, "uid": 1025, "groups": ["analysts"], "username": "oleg.krylov"}, "service": "elasticsearch"}', '2026-05-09 22:00:00'),
(26, '{"host": {"ip": "192.168.1.21", "hostname": "workstation-9"}, "time": {"weekday": "friday", "timestamp": "2026-05-10T08:56:09Z"}, "user": {"gid": 1000, "uid": 1000, "groups": ["clipe-1", "adm", "sudo"], "username": "clipe-1"}, "service": "test"}', '2026-05-10 08:56:09'),
(27, '{"host": {"ip": "192.168.1.22", "hostname": "security-host"}, "time": {"weekday": "saturday", "timestamp": "2026-05-11T13:20:00Z"}, "user": {"gid": 1026, "uid": 1026, "groups": ["security"], "username": "valentina.orlova"}, "service": "systemd"}', '2026-05-11 13:20:00'),
(28, '{"host": {"ip": "172.16.0.34", "hostname": "guest-zone"}, "time": {"weekday": "sunday", "timestamp": "2026-05-12T17:45:00Z"}, "user": {"gid": 1027, "uid": 1027, "groups": ["guests"], "username": "vladislav.ermakov"}, "service": "samba"}', '2026-05-12 17:45:00'),
(29, '{"host": {"ip": "192.168.2.25", "hostname": "cache-2"}, "time": {"weekday": "monday", "timestamp": "2026-05-13T04:15:00Z"}, "user": {"gid": 1028, "uid": 1028, "groups": ["operators"], "username": "alina.titova"}, "service": "redis"}', '2026-05-13 04:15:00'),
(30, '{"host": {"ip": "192.168.1.23", "hostname": "workstation-10"}, "time": {"weekday": "tuesday", "timestamp": "2026-05-14T09:45:00Z"}, "user": {"gid": 1029, "uid": 1029, "groups": ["developers"], "username": "ruslan.mironov"}, "service": "nginx"}', '2026-05-14 09:45:00');

INSERT INTO requests (user_id, context, timestamp) VALUES
(31, '{"host": {"ip": "10.0.0.30", "hostname": "load-balancer"}, "time": {"weekday": "wednesday", "timestamp": "2026-05-15T20:30:00Z"}, "user": {"gid": 1030, "uid": 1030, "groups": ["network"], "username": "vera.belova"}, "service": "nginx"}', '2026-05-15 20:30:00'),
(32, '{"host": {"ip": "192.168.1.24", "hostname": "workstation-11"}, "time": {"weekday": "thursday", "timestamp": "2026-05-16T07:00:00Z"}, "user": {"gid": 1031, "uid": 1031, "groups": ["designers"], "username": "gleb.sergeev"}, "service": "httpd"}', '2026-05-16 07:00:00'),
(33, '{"host": {"ip": "172.16.0.35", "hostname": "vpn-europe"}, "time": {"weekday": "friday", "timestamp": "2026-05-17T12:30:00Z"}, "user": {"gid": 1032, "uid": 1032, "groups": ["remote"], "username": "nadezhda.komarova"}, "service": "ssh"}', '2026-05-17 12:30:00'),
(34, '{"host": {"ip": "192.168.2.26", "hostname": "monitoring-2"}, "time": {"weekday": "saturday", "timestamp": "2026-05-18T23:59:00Z"}, "user": {"gid": 1033, "uid": 1033, "groups": ["monitoring"], "username": "stanislav.medvedev"}, "service": "grafana"}', '2026-05-18 23:59:00'),
(35, '{"host": {"ip": "192.168.1.25", "hostname": "workstation-12"}, "time": {"weekday": "sunday", "timestamp": "2026-05-19T15:30:00Z"}, "user": {"gid": 1034, "uid": 1034, "groups": ["developers", "docker"], "username": "ludmila.borisova"}, "service": "docker"}', '2026-05-19 15:30:00'),
(36, '{"host": {"ip": "10.0.0.31", "hostname": "db-replica"}, "time": {"weekday": "monday", "timestamp": "2026-05-20T03:45:00Z"}, "user": {"gid": 1035, "uid": 1035, "groups": ["dba"], "username": "igor.dmitriev"}, "service": "mysql"}', '2026-05-20 03:45:00'),
(37, '{"host": {"ip": "192.168.1.26", "hostname": "workstation-13"}, "time": {"weekday": "tuesday", "timestamp": "2026-05-21T11:15:00Z"}, "user": {"gid": 1036, "uid": 1036, "groups": ["qa"], "username": "alla.koroleva"}, "service": "jenkins"}', '2026-05-21 11:15:00'),
(38, '{"host": {"ip": "172.16.0.36", "hostname": "guest-network"}, "time": {"weekday": "wednesday", "timestamp": "2026-05-22T18:00:00Z"}, "user": {"gid": 1037, "uid": 1037, "groups": ["guests"], "username": "roman.belousov"}, "service": "ftp"}', '2026-05-22 18:00:00'),
(39, '{"host": {"ip": "192.168.2.27", "hostname": "backup-2"}, "time": {"weekday": "thursday", "timestamp": "2026-05-23T02:30:00Z"}, "user": {"gid": 1038, "uid": 1038, "groups": ["backup"], "username": "inga.terenteva"}, "service": "rsyslog"}', '2026-05-23 02:30:00'),
(40, '{"host": {"ip": "192.168.1.27", "hostname": "workstation-14"}, "time": {"weekday": "friday", "timestamp": "2026-05-24T10:00:00Z"}, "user": {"gid": 1039, "uid": 1039, "groups": ["developers"], "username": "stepan.gromov"}, "service": "gitlab"}', '2026-05-24 10:00:00'),
(41, '{"host": {"ip": "10.0.0.32", "hostname": "cache-3"}, "time": {"weekday": "saturday", "timestamp": "2026-05-25T22:30:00Z"}, "user": {"gid": 1040, "uid": 1040, "groups": ["operators"], "username": "marina.melnikova"}, "service": "redis"}', '2026-05-25 22:30:00'),
(42, '{"host": {"ip": "192.168.1.28", "hostname": "workstation-15"}, "time": {"weekday": "sunday", "timestamp": "2026-05-26T08:30:00Z"}, "user": {"gid": 1041, "uid": 1041, "groups": ["designers"], "username": "vadim.nikonov"}, "service": "httpd"}', '2026-05-26 08:30:00'),
(43, '{"host": {"ip": "172.16.0.37", "hostname": "vpn-asia"}, "time": {"weekday": "monday", "timestamp": "2026-05-27T04:00:00Z"}, "user": {"gid": 1042, "uid": 1042, "groups": ["remote"], "username": "vera.sorokina"}, "service": "ssh"}', '2026-05-27 04:00:00'),
(44, '{"host": {"ip": "192.168.2.28", "hostname": "devops-2"}, "time": {"weekday": "tuesday", "timestamp": "2026-05-28T14:00:00Z"}, "user": {"gid": 1043, "uid": 1043, "groups": ["devops", "sudo"], "username": "gennady.rodionov"}, "service": "kubernetes"}', '2026-05-28 14:00:00'),
(45, '{"host": {"ip": "192.168.1.29", "hostname": "workstation-16"}, "time": {"weekday": "wednesday", "timestamp": "2026-05-29T19:45:00Z"}, "user": {"gid": 1044, "uid": 1044, "groups": ["developers"], "username": "olga.ershova"}, "service": "postgresql"}', '2026-05-29 19:45:00'),
(46, '{"host": {"ip": "10.0.0.33", "hostname": "analytics-2"}, "time": {"weekday": "thursday", "timestamp": "2026-05-30T13:15:00Z"}, "user": {"gid": 1045, "uid": 1045, "groups": ["analysts"], "username": "kirill.zinoviev"}, "service": "kibana"}', '2026-05-30 13:15:00'),
(47, '{"host": {"ip": "192.168.1.30", "hostname": "security-2"}, "time": {"weekday": "friday", "timestamp": "2026-05-31T09:30:00Z"}, "user": {"gid": 1046, "uid": 1046, "groups": ["security"], "username": "lyubov.kalinina"}, "service": "systemd"}', '2026-05-31 09:30:00'),
(48, '{"host": {"ip": "172.16.0.38", "hostname": "wifi-guest"}, "time": {"weekday": "saturday", "timestamp": "2026-06-01T16:00:00Z"}, "user": {"gid": 1047, "uid": 1047, "groups": ["guests"], "username": "artur.bragin"}, "service": "samba"}', '2026-06-01 16:00:00'),
(49, '{"host": {"ip": "192.168.2.29", "hostname": "monitoring-3"}, "time": {"weekday": "sunday", "timestamp": "2026-06-02T21:30:00Z"}, "user": {"gid": 1048, "uid": 1048, "groups": ["monitoring"], "username": "lilia.guseva"}, "service": "prometheus"}', '2026-06-02 21:30:00'),
(50, '{"host": {"ip": "192.168.1.31", "hostname": "workstation-17"}, "time": {"weekday": "monday", "timestamp": "2026-06-03T05:45:00Z"}, "user": {"gid": 1049, "uid": 1049, "groups": ["operators"], "username": "vyacheslav.lazarev"}, "service": "cron"}', '2026-06-03 05:45:00');

INSERT INTO decisions (request_id, policy_id, result, timestamp) VALUES
(1, 1, true, '2026-04-15 09:30:01'), (2, 2, true, '2026-04-16 14:20:01'), (3, 3, false, '2026-04-17 22:15:01'),
(4, 4, true, '2026-04-18 11:45:01'), (5, 5, false, '2026-04-19 16:30:01'), (6, 6, false, '2026-04-20 03:00:01'),
(7, 7, true, '2026-04-21 08:00:01'), (8, 8, true, '2026-04-22 13:30:01'), (9, 9, true, '2026-04-23 23:45:01'),
(10, 10, true, '2026-04-24 10:15:01'), (11, 1, true, '2026-04-25 07:30:01'), (12, 2, true, '2026-04-26 15:00:01'),
(13, 3, false, '2026-04-27 12:00:01'), (14, 4, true, '2026-04-28 18:30:01'), (15, 5, false, '2026-04-29 02:00:01'),
(16, 6, true, '2026-04-30 09:00:01'), (17, 7, false, '2026-05-01 20:15:01'), (18, 8, true, '2026-05-02 01:30:01'),
(19, 9, true, '2026-05-03 14:45:01'), (20, 10, true, '2026-05-04 11:00:01'), (21, 1, false, '2026-05-05 19:00:01'),
(22, 2, true, '2026-05-06 06:00:01'), (23, 3, true, '2026-05-07 16:30:01'), (24, 4, true, '2026-05-08 10:30:01'),
(25, 5, false, '2026-05-09 22:00:01'), (26, 6, true, '2026-05-10 08:56:10'), (27, 7, true, '2026-05-11 13:20:01'),
(28, 8, false, '2026-05-12 17:45:01'), (29, 9, true, '2026-05-13 04:15:01'), (30, 10, true, '2026-05-14 09:45:01'),
(31, 1, false, '2026-05-15 20:30:01'), (32, 2, true, '2026-05-16 07:00:01'), (33, 3, true, '2026-05-17 12:30:01'),
(34, 4, false, '2026-05-18 23:59:01'), (35, 5, true, '2026-05-19 15:30:01'), (36, 6, true, '2026-05-20 03:45:01'),
(37, 7, true, '2026-05-21 11:15:01'), (38, 8, false, '2026-05-22 18:00:01'), (39, 9, true, '2026-05-23 02:30:01'),
(40, 10, true, '2026-05-24 10:00:01'), (41, 1, false, '2026-05-25 22:30:01'), (42, 2, true, '2026-05-26 08:30:01'),
(43, 3, true, '2026-05-27 04:00:01'), (44, 4, true, '2026-05-28 14:00:01'), (45, 5, true, '2026-05-29 19:45:01'),
(46, 6, false, '2026-05-30 13:15:01'), (47, 7, true, '2026-05-31 09:30:01'), (48, 8, false, '2026-06-01 16:00:01'),
(49, 9, true, '2026-06-02 21:30:01'), (50, 10, true, '2026-06-03 05:45:01');