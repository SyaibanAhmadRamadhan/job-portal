CREATE TABLE jobs
(
    id          VARCHAR(255) PRIMARY KEY,
    company_id  VARCHAR(255),
    title       VARCHAR(255) NOT NULL,
    description TEXT         NOT NULL,
    created_at  TIMESTAMP    NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies (id) ON DELETE CASCADE
);
