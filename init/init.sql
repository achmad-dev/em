--- This file is used to create the tables and insert the initial data

--- for storing the users
CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() UNIQUE,
    username text NOT NULL UNIQUE,
    password text NOT NULL,
    role text NOT NULL, -- admin vendor, user hr
    company_name text NOT NULL UNIQUE, -- company name for vendor and hr
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);

--- for storing the events
CREATE TABLE IF NOT EXISTS events (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() UNIQUE,
    user_id uuid NOT NULL,
    vendor_name text NOT NULL, -- vendor name for the event
    event_name text NOT NULL,
    proposed_dates timestamp[],  -- proposed dates for the event
    rejected_remarks text, -- remarks for rejected event
    status text NOT NULL, -- pending, confirmed, rejected
    confirmed_date timestamp, -- date when status is confirmed
    postal_code text NOT NULL, -- postal code of the event location
    location text NOT NULL, -- location of the event
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (vendor_name) REFERENCES users(company_name)
);

--- for logging the user request logs
--- for example, when user request for an event, it will be logged here
CREATE TABLE IF NOT EXISTS request_logs (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() UNIQUE,
    user_id uuid NOT NULL,
    message text NOT NULL,
    status text NOT NULL, -- success, failed
    created_at timestamp with time zone DEFAULT now(),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

--- pre create users
INSERT INTO users (username, password, role, company_name)
SELECT t.username, '$2a$10$dpQJve7Z9/n0rO2EAHbiiOL/ZQl.k7MM7f0xtXkh9OqpS6pQuElBO', t.role, t.company_name
FROM (
    SELECT 
        -- unnest the array to get the username, type and company name
        -- username
        unnest(ARRAY['OneHr', 'TwoHr', 'OneVendor', 'TwoVendor', 'ThreeVendor']) AS username,
        -- role type
        unnest(ARRAY['hr', 'hr', 'vendor', 'vendor', 'vendor']) AS role,
        -- company name
        unnest(ARRAY['OneHr', 'TwoHr', 'OneVendor', 'TwoVendor', 'ThreeVendor']) AS company_name
) t;
