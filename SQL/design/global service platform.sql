-- Table to store user information
DO $$ 
BEGIN 
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_type_enum') THEN
        CREATE TYPE user_type_enum AS ENUM ('client', 'creator');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_tskill_level_enumype_enum') THEN
        CREATE TYPE skill_level_enum AS ENUM ('beginner', 'intermediate', 'advanced');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'demo_type_enum') THEN
        CREATE TYPE demo_type_enum AS ENUM ('image', 'video');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'project_status_enum') THEN
        CREATE TYPE project_status_enum AS ENUM ('pending', 'in_progress', 'done', 'aborted');
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'submission_type_enum') THEN
        CREATE TYPE submission_type_enum AS ENUM ('document', 'code', 'image', 'video');
    END IF;
END $$;

-- Table to store user informationCREATE INDEX IF NOT EXISTS
CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    user_type user_type_enum NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    birth_date DATE,
    profile_image VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS categories (
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for the categories table
CREATE INDEX IF NOT EXISTS idx_categories_category_name ON categories (category_name);

-- Table to store skills associated with categories
CREATE TABLE IF NOT EXISTS skills (
    skill_id SERIAL PRIMARY KEY,
    skill_name VARCHAR(255) NOT NULL UNIQUE,
    category_id INT,
    description TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories(category_id)
);

-- Indexes for the skills table
CREATE INDEX IF NOT EXISTS idx_skills_category_id ON skills (category_id);
CREATE INDEX IF NOT EXISTS idx_skills_skill_name ON skills (skill_name);

-- Table to store cretor skills
CREATE TABLE IF NOT EXISTS creator_skills (
    creator_skill_id SERIAL PRIMARY KEY,
    creator_id INT NOT NULL,
    skill_id INT NOT NULL,
    description TEXT,
    skill_level skill_level_enum NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (creator_id) REFERENCES users(user_id),
    FOREIGN KEY (skill_id) REFERENCES skills(skill_id)
);

-- Indexes for the creator_skills table
CREATE INDEX IF NOT EXISTS idx_creator_skills_creator_id ON creator_skills (creator_id);
CREATE INDEX IF NOT EXISTS idx_creator_skills_skill_id ON creator_skills (skill_id);
CREATE INDEX IF NOT EXISTS idx_creator_skills_skill_level ON creator_skills (skill_level);


-- Table to store demo documents of creators
CREATE TABLE IF NOT EXISTS creator_demos (
    demo_id SERIAL PRIMARY KEY,
    creator_skill_id INT NOT NULL,
    demo_type demo_type_enum NOT NULL,
    demo_path VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (creator_skill_id) REFERENCES creator_skills(creator_skill_id)
);

-- Indexes for the creator_demos table
CREATE INDEX IF NOT EXISTS idx_creator_demos_creator_skill_id ON creator_demos (creator_skill_id);

-- Table to store project details with the current creator assignment
CREATE TABLE IF NOT EXISTS project_details (
    project_id SERIAL PRIMARY KEY,
    project_title VARCHAR(255) NOT NULL,
    client_id INT NOT NULL,
    project_description TEXT NOT NULL, 
    is_locked BOOLEAN DEFAULT FALSE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (client_id) REFERENCES users(user_id)
);

-- Indexes for the project_details table
CREATE INDEX IF NOT EXISTS idx_project_details_title ON project_details (project_title);
CREATE INDEX IF NOT EXISTS idx_project_details_client_id ON project_details (client_id);

CREATE TABLE IF NOT EXISTS raw_materials (
    material_id SERIAL PRIMARY KEY,
    project_id INT NOT NULL,
    material_path VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES project_details(project_id)
);
-- Indexes for the raw_materials table
CREATE INDEX IF NOT EXISTS idx_raw_materials_project_id ON raw_materials (project_id);

CREATE TABLE IF NOT EXISTS project_assignments (
    assignment_id SERIAL PRIMARY KEY,
    project_id INT NOT NULL,
    assigned_creator_id INT NOT NULL,
    deadline DATE NOT NULL,
    project_status project_status_enum DEFAULT 'pending',
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES project_details(project_id),
    FOREIGN KEY (assigned_creator_id) REFERENCES users(user_id)
);

-- Indexes for the project_assignments table
CREATE INDEX IF NOT EXISTS idx_project_assignments_project_status ON project_assignments (project_status);
CREATE INDEX IF NOT EXISTS idx_project_assignments_project_id ON project_assignments (project_id);
CREATE INDEX IF NOT EXISTS idx_project_assignments_assigned_creator_id ON project_assignments (assigned_creator_id);

-- Table to store ongoing communication between clients and creators
CREATE TABLE IF NOT EXISTS project_communication (
    communication_id SERIAL PRIMARY KEY,
    project_id INT,
    sender_user_id INT NOT NULL,
    receiver_user_id INT NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (project_id) REFERENCES project_details(project_id),
    FOREIGN KEY (sender_user_id) REFERENCES users(user_id),
    FOREIGN KEY (receiver_user_id) REFERENCES users(user_id)
);

-- Indexes for the project_communication table
CREATE INDEX IF NOT EXISTS idx_project_communication_project_id ON project_communication (project_id);
CREATE INDEX IF NOT EXISTS idx_project_communication_sender_user_id ON project_communication (sender_user_id);
CREATE INDEX IF NOT EXISTS idx_project_communication_receiver_user_id ON project_communication (receiver_user_id);

-- Table to store task submissions by creators
CREATE TABLE IF NOT EXISTS task_submissions (
    submission_id SERIAL PRIMARY KEY,
    assignment_id INT,
    submission_type submission_type_enum DEFAULT 'document',
    submission_path VARCHAR(255) NOT NULL,
    submission_description TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (assignment_id) REFERENCES project_assignments(assignment_id)
);

-- Indexes for the task_submissions table
CREATE INDEX IF NOT EXISTS idx_task_submissions_assignment_id ON task_submissions (assignment_id);
CREATE INDEX IF NOT EXISTS idx_task_submissions_submission_type ON task_submissions (submission_type);

-- Table to store feedback from clients
CREATE TABLE IF NOT EXISTS client_feedback (
    feedback_id SERIAL PRIMARY KEY,
    submission_id INT,
    client_id INT,
    feedback_text TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (submission_id) REFERENCES task_submissions(submission_id),
    FOREIGN KEY (client_id) REFERENCES users(user_id)
);

-- Indexes for the client_feedback table
CREATE INDEX IF NOT EXISTS idx_client_feedback_submission_id ON client_feedback (submission_id);
CREATE INDEX IF NOT EXISTS idx_client_feedback_client_id ON client_feedback (client_id);