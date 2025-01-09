CREATE TABLE jobs
(
    id          VARCHAR(255) PRIMARY KEY,
    company_id  VARCHAR(255),
    title       VARCHAR(255) NOT NULL,
    description TEXT         NOT NULL,
    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE
);
