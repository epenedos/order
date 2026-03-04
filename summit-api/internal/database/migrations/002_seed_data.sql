-- Summit Sporting Goods - Seed Data
-- Equivalent to the Oracle summit.dmp data, adapted for PostgreSQL
-- Run after 001_initial_schema.sql

BEGIN;

-- ============================================================
-- REGIONS (from S_REGION)
-- ============================================================
INSERT INTO regions (id, name) VALUES
(1, 'North America'),
(2, 'South America'),
(3, 'Africa / Middle East'),
(4, 'Asia'),
(5, 'Europe');
SELECT setval('regions_id_seq', 5);

-- ============================================================
-- TITLES (from S_TITLE - job title lookup)
-- ============================================================
INSERT INTO titles (title) VALUES
('President'),
('VP, Operations'),
('VP, Sales'),
('VP, Finance'),
('VP, Administration'),
('Director, Operations'),
('Director, Sales'),
('Director, Finance'),
('Director, Administration'),
('Warehouse Manager'),
('Sales Representative'),
('Stock Clerk');

-- ============================================================
-- DEPARTMENTS (from S_DEPT)
-- ============================================================
INSERT INTO departments (id, name, region_id) VALUES
(10, 'Finance',        1),
(31, 'Sales',          1),
(32, 'Sales',          2),
(33, 'Sales',          3),
(34, 'Sales',          4),
(35, 'Sales',          5),
(41, 'Operations',     1),
(42, 'Operations',     2),
(43, 'Operations',     3),
(44, 'Operations',     4),
(45, 'Operations',     5),
(50, 'Administration', 1);
SELECT setval('departments_id_seq', 50);

-- ============================================================
-- EMPLOYEES (from S_EMP)
-- ============================================================
INSERT INTO employees (id, last_name, first_name, userid, start_date, comments, manager_id, title, dept_id, salary, commission_pct) VALUES
(1,  'Velasquez',  'Carmen',   'cvelasqu', '1990-03-03', NULL,                          NULL, 'President',            50, 2500.00, NULL),
(2,  'Ngao',       'LaDoris',  'lngao',    '1990-03-08', NULL,                          1,    'VP, Operations',       41, 1450.00, NULL),
(3,  'Nagayama',   'Midori',   'mnagayam', '1991-06-17', NULL,                          1,    'VP, Sales',            31, 1400.00, NULL),
(4,  'Quick-To-See','Mark',    'mquickto', '1990-04-07', NULL,                          1,    'VP, Finance',          10, 1450.00, NULL),
(5,  'Ropeburn',   'Audry',    'aropebur', '1990-03-04', NULL,                          1,    'VP, Administration',   50, 1550.00, NULL),
(6,  'Urguhart',   'Molly',    'murguhar', '1991-01-18', NULL,                          2,    'Warehouse Manager',    41, 1200.00, NULL),
(7,  'Menchu',     'Roberta',  'rmenchu',  '1990-05-14', NULL,                          2,    'Warehouse Manager',    42, 1250.00, NULL),
(8,  'Biri',       'Ben',      'bbiri',    '1990-04-07', NULL,                          2,    'Warehouse Manager',    43, 1100.00, NULL),
(9,  'Catchpole',  'Antoinette','acatchpo', '1992-02-09', NULL,                         2,    'Warehouse Manager',    44, 1300.00, NULL),
(10, 'Havel',      'Marta',    'mhavel',   '1991-02-27', NULL,                          2,    'Warehouse Manager',    45, 1307.00, NULL),
(11, 'Magee',      'Colin',    'cmagee',   '1990-05-14', NULL,                          3,    'Sales Representative', 31, 1400.00, 10.00),
(12, 'Giljum',     'Henry',    'hgiljum',  '1992-01-18', NULL,                          3,    'Sales Representative', 32, 1490.00, 12.50),
(13, 'Sedeghi',    'Yasmin',   'ysedeghi', '1991-02-18', NULL,                          3,    'Sales Representative', 33, 1515.00, 10.00),
(14, 'Nguyen',     'Mai',      'mnguyen',  '1992-01-22', NULL,                          3,    'Sales Representative', 34, 1525.00, 15.00),
(15, 'Dumas',      'Andre',    'adumas',   '1991-10-09', NULL,                          3,    'Sales Representative', 35, 1450.00, 17.50),
(16, 'Maduro',     'Elena',    'emaduro',  '1992-02-07', 'Fluent in French and Spanish', 6,    'Stock Clerk',          41,  1400.00, NULL),
(17, 'Smith',      'George',   'gsmith',   '1990-03-08', NULL,                          6,    'Stock Clerk',          41,  940.00, NULL),
(18, 'Nozaki',     'Akira',    'anozaki',  '1991-02-09', NULL,                          7,    'Stock Clerk',          42, 1200.00, NULL),
(19, 'Patel',      'Vikram',   'vpatel',   '1991-08-06', NULL,                          7,    'Stock Clerk',          42, 795.00,  NULL),
(20, 'Newman',     'Chad',     'cnewman',  '1991-07-21', NULL,                          8,    'Stock Clerk',          43, 750.00,  NULL),
(21, 'Markarian',  'Alexander','amarkari', '1991-05-26', NULL,                          8,    'Stock Clerk',          43, 850.00,  NULL),
(22, 'Chang',      'Eddie',    'echang',   '1990-11-30', NULL,                          9,    'Stock Clerk',          44, 800.00,  NULL),
(23, 'Patel',      'Radha',    'rpatel',   '1990-10-17', NULL,                          9,    'Stock Clerk',          44, 795.00,  NULL),
(24, 'Dancs',      'Bela',     'bdancs',   '1991-03-17', NULL,                          10,   'Stock Clerk',          45, 860.00,  NULL),
(25, 'Schwartz',   'Sylvie',   'sschwart', '1991-05-09', NULL,                          10,   'Stock Clerk',          45, 1100.00, NULL);
SELECT setval('employees_id_seq', 25);

-- ============================================================
-- CUSTOMERS (from S_CUSTOMER)
-- ============================================================
INSERT INTO customers (id, name, phone, address, city, state, country, zip_code, credit_rating, sales_rep_id, region_id, comments) VALUES
(201, 'Unisports',           '55-2066101',   '72 Via Bahia',                  'Sao Paolo',       NULL,    'Brazil',          NULL,    'EXCELLENT', 12, 2, NULL),
(202, 'OJ Athletics',        '81-20101',     '6741 Takashi Blvd.',            'Osaka',            NULL,    'Japan',           NULL,    'POOR',      14, 4, NULL),
(203, 'Delhi Sports',        '91-10351',     '11368 Chanakya',                'New Delhi',        NULL,    'India',           NULL,    'GOOD',      14, 4, NULL),
(204, 'Womansport',          '1-206-104-0103','281 King Street',              'Seattle',          'WA',    'United States',   NULL,    'EXCELLENT', 11, 1, NULL),
(205, 'Kam''s Sporting Goods','852-3692888',  '15 Henessey Road',             'Hong Kong',        NULL,    'China',           NULL,    'EXCELLENT', 15, 4, NULL),
(206, 'Sportique',           '33-2257201',   '172 Rue de Rivoli',            'Cannes',           NULL,    'France',          NULL,    'EXCELLENT', 15, 5, NULL),
(207, 'Sweet Rock Sports',   '234-6036201',  '6 Saint Antoine',              'Lagos',            NULL,    'Nigeria',         NULL,    'GOOD',      13, 3, NULL),
(208, 'Muench Sports',       '49-527454',    '435 Grunerstrasse',            'Stuttgart',        NULL,    'Germany',         NULL,    'GOOD',      15, 5, NULL),
(209, 'Beisbol Si!',         '809-352689',   '792 Playa Del Mar',            'San Pedro de Macoris', NULL,'Dominican Republic', NULL, 'EXCELLENT', 11, 1, NULL),
(210, 'Futbol Sonora',       '52-404562',    '3 Via Saguaro',                'Nogales',          NULL,    'Mexico',          NULL,    'EXCELLENT', 12, 2, NULL),
(211, 'Kuhn''s Sports',      '42-111292',    '7 Modrany',                    'Prague',           NULL,    'Czechoslovakia',  NULL,    'EXCELLENT', 15, 5, NULL),
(212, 'Hamada Sport',        '20-1209211',   '57A Corniche',                 'Alexandria',       NULL,    'Egypt',           NULL,    'EXCELLENT', 13, 3, NULL),
(213, 'Big John''s Sports Emporium', '1-415-555-6281', '4783 18th Street',  'San Francisco',    'CA',    'United States',   NULL,    'EXCELLENT', 11, 1, NULL),
(214, 'Ojibway Retail',      '1-716-555-7171','415 Main Street',             'Buffalo',          'NY',    'United States',   NULL,    'POOR',      11, 1, NULL),
(215, 'Sporta Russia',       '7-3892456',    '6000 Yekaterinskaya',          'Saint Petersburg',  NULL,   'Russia',          NULL,    'POOR',      15, 5, NULL);
SELECT setval('customers_id_seq', 215);

-- ============================================================
-- LONG_TEXTS (from S_LONGTEXT — product descriptions)
-- ============================================================
INSERT INTO long_texts (id, use_filename, filename, text_content) VALUES
(518,  FALSE, NULL, 'Protective equipment for the serious and casual player.'),
(519,  FALSE, NULL, 'Bats and balls for any level of play.'),
(520,  FALSE, NULL, 'Comfort and style for every level and every court.'),
(527,  FALSE, NULL, 'High-technology running shoes for increased speed and comfort.'),
(528,  FALSE, NULL, 'Protective headgear for the player who takes safety seriously.'),
(529,  FALSE, NULL, 'Premium quality tennis racquets for every level of play.'),
(530,  FALSE, NULL, 'A fine selection of tennis balls at a reasonable price.'),
(531,  FALSE, NULL, 'Sturdy, well-designed frames for recreational and serious cyclists.'),
(532,  FALSE, NULL, 'Aerodynamic wheels designed for speed and endurance.'),
(533,  FALSE, NULL, 'Top-of-the-line saddles designed to make those long rides more comfortable.'),
(534,  FALSE, NULL, 'A combination of style and comfort that cannot be beat.'),
(535,  FALSE, NULL, 'Maximum protection for the daring downhill and cross-country skier.'),
(536,  FALSE, NULL, 'The ultimate in downhill racing skis.'),
(537,  FALSE, NULL, 'Top-of-the-line poles for every level of skier.'),
(538,  FALSE, NULL, 'State-of-the-art ski boots for all levels.'),
(539,  FALSE, NULL, 'A complete range of goggles for both sun and snow protection.');
SELECT setval('long_texts_id_seq', 539);

-- ============================================================
-- IMAGES (from S_IMAGE — product image references)
-- ============================================================
INSERT INTO images (id, format, use_filename, filename, image_data) VALUES
(1001, 'TIFF', TRUE, 'baseball.tif',     NULL),
(1002, 'TIFF', TRUE, 'batting.tif',      NULL),
(1003, 'TIFF', TRUE, 'shoes.tif',        NULL),
(1004, 'TIFF', TRUE, 'running.tif',      NULL),
(1005, 'TIFF', TRUE, 'helmet.tif',       NULL),
(1006, 'TIFF', TRUE, 'racquet.tif',      NULL),
(1007, 'TIFF', TRUE, 'tennisball.tif',   NULL),
(1008, 'TIFF', TRUE, 'bicycle.tif',      NULL),
(1009, 'TIFF', TRUE, 'wheel.tif',        NULL),
(1010, 'TIFF', TRUE, 'saddle.tif',       NULL),
(1011, 'TIFF', TRUE, 'sunglasses.tif',   NULL),
(1012, 'TIFF', TRUE, 'skiboot.tif',      NULL),
(1013, 'TIFF', TRUE, 'ski.tif',          NULL),
(1014, 'TIFF', TRUE, 'skipole.tif',      NULL),
(1015, 'TIFF', TRUE, 'skiboot2.tif',     NULL),
(1016, 'TIFF', TRUE, 'goggle.tif',       NULL);
SELECT setval('images_id_seq', 1016);

-- ============================================================
-- PRODUCTS (from S_PRODUCT)
-- ============================================================
INSERT INTO products (id, name, short_desc, longtext_id, image_id, suggested_whlsl_price, whlsl_units) VALUES
(10011, 'Bunny Boot',           'Ski boot  - Loss leader',         538, 1015,  150.00, NULL),
(10012, 'Ace Ski Boot',         'Ski boot  - Loss leader',         538, 1012,  200.00, NULL),
(10013, 'Pro Ski Boot',         'Ski boot  - Moderate markup',     538, 1012,  410.00, NULL),
(10021, 'Bunny Ski Pole',       'Ski pole  - Low-end',             537, 1014,   16.25, NULL),
(10022, 'Ace Ski Pole',         'Ski pole  - Mid-range',           537, 1014,   21.95, NULL),
(10023, 'Pro Ski Pole',         'Ski pole  - Top-of-line',         537, 1014,   40.95, NULL),
(20106, 'Junior Soccer Ball',   'Junior size  - Loss leader',      NULL, NULL,   11.00, NULL),
(20108, 'World Cup Soccer Ball','Official size and weight',        NULL, NULL,   28.00, NULL),
(20201, 'World Cup Net',        'Official FIFA-approved',          NULL, NULL,  123.00, NULL),
(20301, 'World Cup Shin Pads',  'Designed for world-class play',   NULL, NULL,   36.00, NULL),
(20510, 'Black Hawk    Bat',   'Graphite bat  - Loss leader',     519, 1002,  240.00, NULL),
(20512, 'Bat Pack',             'Complete bat kit  - Moderate markup', 519, 1002, 55.00, NULL),
(30321, 'Grand Prix Bicycle',   'Top-of-line racing bicycle',      531, 1008, 1669.00, NULL),
(30326, 'Himalaya Bicycle',     'Mountain bicycle',                531, 1008,  582.00, NULL),
(30421, 'Grand Prix Bicycle Tires','Aerodynamic tubular tires',    532, 1009,   16.00, NULL),
(30426, 'Himalaya Tires',       'Knobby mountain bike tires',      532, 1009,   18.25, NULL),
(30433, 'New Air Pump',         'Portable air pump',               NULL, NULL,   20.00, NULL),
(32779, 'Ace Tennis Racquet',   'Graphite frame  - Loss leader',   529, 1006,   58.00, NULL),
(32861, 'Ace Tennis Balls',     'High-quality 6-ball can',         530, 1007,    3.00, NULL),
(40421, 'Alexeyer Pro Lifting Bar','Chrome bar',                   NULL, NULL,   65.00, NULL),
(40422, 'Pro Curling Bar',      'Curling bar  - Moderate markup',  NULL, NULL,   50.00, NULL),
(41010, 'Bunny Gloves',        'Beginner  - Loss leader',          518, 1001,   14.00, NULL),
(41020, 'Ace Gloves',          'Advanced  - Moderate markup',      518, 1001,   19.00, NULL),
(41050, 'Pro Gloves',          'Professional quality',             518, 1001,   28.00, NULL),
(41080, 'Pro Batting Helmet',  'Professional batting helmet',      528, 1005,   60.00, NULL),
(41100, 'Major League Baseball','Official major league',           519, 1002,    4.29, NULL),
(50169, 'Championship Basketball','Professional quality',          520, 1003,   43.00, NULL),
(50273, 'Pro Skateboard',      'Titanium trucks',                  NULL, NULL,   33.00, NULL),
(50417, 'Leather Basketball',  'High-quality leather',             520, 1003,   32.00, NULL),
(50418, 'Rubber Basketball',   'High-quality rubber',              520, 1003,   25.00, NULL),
(50530, 'Ace Racer',           'High-performance running shoes',   527, 1004,   75.00, NULL),
(50532, 'Pro Runner',          'Professional running shoes',       527, 1004,   95.00, NULL),
(50536, 'Air Sneaker',         'Casual athletic shoe',             527, 1003,   65.00, NULL),
(50537, 'Mohawk Running Shoe', 'Running shoe  - Loss leader',      527, 1004,   42.00, NULL);
SELECT setval('products_id_seq', 50537);

-- ============================================================
-- WAREHOUSES (from S_WAREHOUSE)
-- ============================================================
INSERT INTO warehouses (id, region_id, address, city, state, country, zip_code, phone, manager_id) VALUES
(101, 1, '283 King Street',     'Seattle',          'WA',  'United States', NULL, NULL, 6),
(10501, 2, '5765 N. 10th Street', 'Sao Paolo',      NULL,  'Brazil',        NULL, NULL, 7),
(201, 3, '68 Via Centrale',      'Sao Paolo',        NULL,  'Brazil',        NULL, NULL, 7),
(301, 3, '6921 King Way',        'Lagos',            NULL,  'Nigeria',       NULL, NULL, 8),
(401, 4, '86 Chu Street',        'Hong Kong',        NULL,  'China',         NULL, NULL, 9),
(501, 5, '5765 N. 10th Street',  'Belmont',          NULL,  'France',        NULL, NULL, 10);
SELECT setval('warehouses_id_seq', 10501);

-- ============================================================
-- INVENTORY (from S_INVENTORY)
-- ============================================================
INSERT INTO inventory (product_id, warehouse_id, amount_in_stock, reorder_point, max_in_stock, out_of_stock_explanation, restock_date) VALUES
(10011, 101,   650, 625, 1100, NULL, NULL),
(10012, 101,   600, 560, 1000, NULL, NULL),
(10013, 101,   400, 400,  700, NULL, NULL),
(10021, 101,   500, 425,  740, NULL, NULL),
(10022, 101,   300, 200,  350, NULL, NULL),
(10023, 101,   400, 300,  525, NULL, NULL),
(20106, 101,   993, 625, 1000, NULL, NULL),
(20108, 101,   700, 700, 1225, NULL, NULL),
(20201, 101,   802, 800, 1400, NULL, NULL),
(20301, 101,   500, 500,  875, NULL, NULL),
(20510, 101,   100, 200,  350, 'Manufacturer backlog', '1993-08-06'),
(20512, 101,   300, 300,  525, NULL, NULL),
(30321, 101,   300, 300,  525, NULL, NULL),
(30326, 101,   600, 560, 1000, NULL, NULL),
(30421, 101,   400, 400,  700, NULL, NULL),
(30426, 101,   500, 425,  740, NULL, NULL),
(30433, 101,   300, 300,  525, NULL, NULL),
(32779, 101,   500, 425,  740, NULL, NULL),
(32861, 101,   600, 560, 1000, NULL, NULL),
(40421, 101,   750, 600, 1050, NULL, NULL),
(40422, 101,   750, 600, 1050, NULL, NULL),
(41010, 101,   500, 425,  740, NULL, NULL),
(41020, 101,   500, 425,  740, NULL, NULL),
(41050, 101,   400, 400,  700, NULL, NULL),
(41080, 101,   400, 400,  700, NULL, NULL),
(41100, 101,   500, 425,  740, NULL, NULL),
(50169, 101,   400, 400,  700, NULL, NULL),
(50273, 101,   500, 425,  740, NULL, NULL),
(50417, 101,   400, 300,  525, NULL, NULL),
(50418, 101,   500, 425,  740, NULL, NULL),
(50530, 101,   400, 400,  700, NULL, NULL),
(50532, 101,   300, 300,  525, NULL, NULL),
(50536, 101,   600, 560, 1000, NULL, NULL),
(50537, 101,   400, 400,  700, NULL, NULL),
-- Warehouse 10501 (Sao Paolo)
(10012, 10501, 200, 150,  300, NULL, NULL),
(10022, 10501, 300, 200,  350, NULL, NULL),
(20108, 10501, 150, 100,  175, NULL, NULL),
(20201, 10501, 200, 150,  275, NULL, NULL),
-- Warehouse 201 (Brazil)
(10011, 201,   250, 200,  350, NULL, NULL),
(10012, 201,   300, 250,  450, NULL, NULL),
(10022, 201,   280, 250,  440, NULL, NULL),
(20106, 201,   100, 100,  175, NULL, NULL),
(20108, 201,   200, 150,  275, NULL, NULL),
(20510, 201,    50,  75,  131, 'Seasonal', NULL),
(32779, 201,   150, 150,  263, NULL, NULL),
(32861, 201,   200, 150,  275, NULL, NULL),
(50169, 201,   200, 150,  275, NULL, NULL),
(50530, 201,   250, 200,  350, NULL, NULL),
(50532, 201,   100, 100,  175, NULL, NULL),
-- Warehouse 301 (Lagos)
(10012, 301,   100,  75,  131, NULL, NULL),
(10023, 301,   200, 150,  275, NULL, NULL),
(20106, 301,   150, 100,  175, NULL, NULL),
(20201, 301,   100, 100,  175, NULL, NULL),
(20301, 301,   120, 100,  175, NULL, NULL),
(30321, 301,   100, 100,  175, NULL, NULL),
(41010, 301,   200, 150,  275, NULL, NULL),
(41100, 301,   250, 200,  350, NULL, NULL),
(50417, 301,   200, 150,  275, NULL, NULL),
(50418, 301,   250, 200,  350, NULL, NULL),
-- Warehouse 401 (Hong Kong)
(10011, 401,   300, 250,  440, NULL, NULL),
(10012, 401,   350, 300,  525, NULL, NULL),
(10022, 401,   200, 200,  350, NULL, NULL),
(20510, 401,   150, 100,  175, NULL, NULL),
(30321, 401,   150, 100,  175, NULL, NULL),
(30326, 401,   200, 200,  350, NULL, NULL),
(30421, 401,   100, 100,  175, NULL, NULL),
(41020, 401,   200, 150,  275, NULL, NULL),
(41080, 401,   200, 150,  275, NULL, NULL),
(50273, 401,   200, 200,  350, NULL, NULL),
(50536, 401,   200, 200,  350, NULL, NULL),
(50537, 401,   150, 100,  175, NULL, NULL),
-- Warehouse 501 (France)
(10011, 501,   300, 200,  350, NULL, NULL),
(10013, 501,   200, 200,  350, NULL, NULL),
(10021, 501,   200, 150,  275, NULL, NULL),
(10023, 501,   300, 200,  350, NULL, NULL),
(20106, 501,   200, 200,  350, NULL, NULL),
(20108, 501,   250, 200,  350, NULL, NULL),
(20301, 501,   300, 200,  350, NULL, NULL),
(30321, 501,   200, 150,  275, NULL, NULL),
(30421, 501,   200, 150,  275, NULL, NULL),
(30433, 501,   100, 100,  175, NULL, NULL),
(32779, 501,   200, 200,  350, NULL, NULL),
(40421, 501,   250, 200,  350, NULL, NULL),
(40422, 501,   250, 200,  350, NULL, NULL),
(41010, 501,   200, 200,  350, NULL, NULL),
(41050, 501,   200, 150,  275, NULL, NULL),
(50169, 501,   250, 200,  350, NULL, NULL),
(50530, 501,   250, 200,  350, NULL, NULL),
(50532, 501,   200, 200,  350, NULL, NULL),
(50536, 501,   250, 200,  350, NULL, NULL);

-- ============================================================
-- ORDERS (from S_ORD)
-- ============================================================
INSERT INTO orders (id, customer_id, date_ordered, date_shipped, sales_rep_id, total, payment_type, order_filled) VALUES
(100, 204, '1992-08-31', '1992-09-10', 11,  601100.00, 'CREDIT', TRUE),
(101, 205, '1992-08-31', '1992-09-15', 14,    8056.60, 'CREDIT', TRUE),
(102, 206, '1992-09-01', '1992-09-08', 15,    8335.00, 'CREDIT', TRUE),
(103, 208, '1992-09-02', '1992-09-22', 15,     377.00, 'CASH',   TRUE),
(104, 208, '1992-09-03', '1992-09-23', 15,   32430.00, 'CREDIT', TRUE),
(105, 209, '1992-09-04', '1992-09-18', 11,    2722.24, 'CREDIT', TRUE),
(106, 210, '1992-09-07', '1992-09-15', 12,   15634.00, 'CREDIT', TRUE),
(107, 211, '1992-09-07', '1992-09-21', 15,     142.50, 'CREDIT', TRUE),
(108, 212, '1992-09-07', '1992-09-10', 13,  149570.00, 'CREDIT', TRUE),
(109, 213, '1992-09-08', '1992-09-28', 11,   10794.60, 'CREDIT', FALSE),
(110, 214, '1992-09-09', NULL,          11,    1539.13, 'CASH',   FALSE),
(111, 204, '1992-09-09', NULL,          11,    2770.00, 'CASH',   FALSE),
(97,  201, '1992-08-28', '1992-09-17', 12,   84000.00, 'CREDIT', TRUE),
(98,  202, '1992-08-31', '1992-09-10', 14,     595.00, 'CASH',   TRUE),
(99,  203, '1992-08-31', '1992-09-18', 14,    7707.00, 'CREDIT', TRUE);
SELECT setval('orders_id_seq', 111);

-- ============================================================
-- ORDER_ITEMS (from S_ITEM)
-- ============================================================
INSERT INTO order_items (ord_id, item_id, product_id, price, quantity, quantity_shipped) VALUES
-- Order 97
(97,  1, 20106,   9.00,  1000,  1000),
(97,  2, 30321,  75.00,  1000,  1000),
-- Order 98
(98,  1, 40421,  85.00,     7,     7),
-- Order 99
(99,  1, 20510, 350.00,    20,    20),
(99,  2, 40421,  97.00,     3,     3),
(99,  3, 41100,   6.00,  1000,  1000),
-- Order 100
(100, 1, 30421, 16.00, 500, 500),
(100, 2, 30326, 600.00, 1000, 1000),
(100, 3, 50536, 60.00, 18, 18),
-- Order 101
(101, 1, 30421,  15.00,   8,     8),
(101, 2, 41010,  14.00,   3,     3),
(101, 3, 50169,  43.00,   5,     5),
(101, 4, 50417,  30.00,   6,     6),
(101, 5, 50530,  79.00,  10,    10),
(101, 6, 50536,  60.00,  12,    12),
(101, 7, 50537,  40.00,  10,    10),
(101, 8, 10013, 350.00,   4,     4),
(101, 9, 10022,  21.95,  10,    10),
(101, 10, 10023, 36.00,  10,    10),
-- Order 102
(102, 1, 20108,  28.00, 100, 100),
(102, 2, 20201, 115.00,  45,  45),
(102, 3, 50530,  79.00,  10,  10),
(102, 4, 50536,  62.00,  10,  10),
-- Order 103
(103, 1, 30433, 20.00, 1, 1),
(103, 2, 32861,  3.00, 5, 5),
(103, 3, 50417, 32.00, 1, 1),
(103, 4, 50418, 25.00, 1, 1),
(103, 5, 50537, 40.00, 7, 7),
-- Order 104
(104, 1, 20510, 350.00,   10,  10),
(104, 2, 20512,  55.00,   50,  50),
(104, 3, 30321,1669.00,    3,   3),
(104, 4, 30326, 582.00,   32,  32),
(104, 5, 32779,  58.00,   50,  50),
-- Order 105
(105, 1, 50530,  79.00,   10,  10),
(105, 2, 50536,  60.00,   12,  12),
(105, 3, 50537,  42.00,   18,  18),
(105, 4, 20108,  28.00,    3,   3),
(105, 5, 20201, 115.00,    3,   3),
-- Order 106
(106, 1, 20108,  28.00,  100, 100),
(106, 2, 20201, 123.00,   45,  45),
(106, 3, 30421,  16.00,  100, 100),
(106, 4, 30421,  16.00,  100, 100),
(106, 5, 30433,  20.00,   50,  50),
-- Order 107
(107, 1, 20106,   9.00,    5,   5),
(107, 2, 20108,  28.00,    2,   2),
(107, 3, 30433,  20.00,    1,   1),
-- Order 108
(108, 1, 20510, 350.00,  250, 250),
(108, 2, 30326, 582.00,   60,  60),
(108, 3, 41080,  35.00,  200, 200),
(108, 4, 41100,   5.00, 1000,1000),
(108, 5, 50273,  33.00,  500, 500),
(108, 6, 50418,  25.00,  500, 500),
(108, 7, 50537,  42.00,  500, 500),
-- Order 109
(109, 1, 10011, 140.00,   50,   0),
(109, 2, 10013, 390.00,   12,   0),
(109, 3, 10021,  16.25,   24,   0),
(109, 4, 10023,  36.00,   12,   0),
(109, 5, 30421,  16.00,  100,   0),
(109, 6, 41010,  14.00,   30,   0),
(109, 7, 50530,  79.00,   18,   0),
-- Order 110
(110, 1, 50536,  60.00,   10,   0),
(110, 2, 50537,  42.00,   10,   0),
(110, 3, 20510, 350.00,    1,   0),
(110, 4, 20512,  55.00,    3,   0),
-- Order 111
(111, 1, 20106,   11.00,  50,   0),
(111, 2, 30421,   16.00, 100,   0),
(111, 3, 32779,   58.00,   2,   0);

-- ============================================================
-- USERS (app auth — seeded admin + sales reps)
-- ============================================================
-- password for all seeded users: "summit123"
-- bcrypt hash of "summit123"
INSERT INTO users (id, email, password_hash, employee_id, role, is_active) VALUES
(1, 'admin@summit.com',    '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 1,  'admin',     TRUE),
(2, 'cmagee@summit.com',   '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 11, 'sales_rep', TRUE),
(3, 'hgiljum@summit.com',  '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 12, 'sales_rep', TRUE),
(4, 'ysedeghi@summit.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 13, 'sales_rep', TRUE),
(5, 'mnguyen@summit.com',  '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 14, 'sales_rep', TRUE),
(6, 'adumas@summit.com',   '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 15, 'sales_rep', TRUE),
(7, 'viewer@summit.com',   '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', NULL, 'viewer',  TRUE);
SELECT setval('users_id_seq', 7);

COMMIT;
