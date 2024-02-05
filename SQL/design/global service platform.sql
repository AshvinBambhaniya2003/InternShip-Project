-- Table to store user information
DO $$ 
BEGIN 
    -- Table: users
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'users') THEN 
       CREATE TYPE user_type_enum AS ENUM ('client', 'creator'); 
       CREATE TABLE users (
            user_id SERIAL PRIMARY KEY,
            username VARCHAR(255) NOT NULL,
            email VARCHAR(255) UNIQUE NOT NULL,
            user_type user_type_enum NOT NULL,
            first_name VARCHAR(100),
            last_name VARCHAR(100),
            birth_date DATE,
            profile_image VARCHAR(255),
            created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
        );

        -- Indexes for the user table
        CREATE INDEX idx_user_username ON users(username);
        CREATE INDEX idx_user_email ON users(email);
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'creators') THEN 
        -- Table to store creator profiles
        CREATE TABLE creators (
            creator_id SERIAL PRIMARY KEY,
            user_id INT UNIQUE NOT NULL,
            description TEXT,
            FOREIGN KEY (user_id) REFERENCES users(user_id)
        );

        -- Indexes for the creators table
        CREATE INDEX idx_creators_user_id ON creators(user_id); 
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'skill_categories') THEN 
        -- Table to store categories of skills
        CREATE TABLE skill_categories (
            category_id SERIAL PRIMARY KEY,
            category_name VARCHAR(255) NOT NULL UNIQUE,
            created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
        );
    END IF;
    
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'skills') THEN 
        -- Table to store skills associated with categories
        CREATE TYPE skill_level_enum AS ENUM ('beginner', 'intermediate', 'advanced');
        CREATE TABLE skills (
            skill_id SERIAL PRIMARY KEY,
            skill_name VARCHAR(255) NOT NULL UNIQUE,
            category_id INT,
            description TEXT,
            skill_level skill_level_enum NOT NULL,
            created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (category_id) REFERENCES skill_categories(category_id)
        );

        -- Indexes for the skills table
        CREATE INDEX idx_skills_skill_name ON skills(skill_name);
        CREATE INDEX idx_skills_category_id ON skills(category_id);
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'creator_skills') THEN 
        -- Table to store the relationship between creators and skills
        CREATE TABLE creator_skills (
            creator_skill_id SERIAL PRIMARY KEY,
            creator_id INT NOT NULL,
            skill_id INT NOT NULL,
            experience_years INT,
            FOREIGN KEY (creator_id) REFERENCES creators(creator_id),
            FOREIGN KEY (skill_id) REFERENCES skills(skill_id)
        );

        -- Indexes for the creator_skills table
        CREATE INDEX idx_creator_skills_creator_id ON creator_skills(creator_id);
        CREATE INDEX idx_creator_skills_skill_id ON creator_skills(skill_id);
    END IF;
    
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'creator_demos') THEN 
        -- Table to store demo documents of creators

        CREATE TYPE document_type_enum AS ENUM ('image', 'video');
        CREATE TABLE creator_demos (
            demo_id SERIAL PRIMARY KEY,
            creator_skill_id INT NOT NULL,
            document_type document_type_enum NOT NULL,
            document_path VARCHAR(255) NOT NULL,
            created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (creator_skill_id) REFERENCES creator_skills(creator_skill_id)
        );

        -- Indexes for the creator_demos table
        CREATE INDEX idx_creator_demos_creator_skill_id ON creator_demos(creator_skill_id);
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'projects') THEN 
        -- Table to store project information

        CREATE TYPE project_status_enum AS ENUM ('pending', 'in_progress', 'done', 'aborted');
        CREATE TABLE projects (
            project_id SERIAL PRIMARY KEY,
            client_id INT NOT NULL,
            project_title VARCHAR(255) NOT NULL,
            deadline DATE NOT NULL,
            assigned_creator_id INT,
            project_status project_status_enum DEFAULT 'pending',
            created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (client_id) REFERENCES users(user_id),
            FOREIGN KEY (assigned_creator_id) REFERENCES creators(creator_id)
        );

        -- Indexes for the projects table
        CREATE INDEX idx_projects_client_id ON projects(client_id);
        CREATE INDEX idx_projects_assigned_creator_id ON projects(assigned_creator_id);
        CREATE INDEX idx_projects_status ON projects(project_status);
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'project_details') THEN 
        -- Table to store project details with the current creator assignment
        CREATE TABLE project_details (
            project_detail_id SERIAL PRIMARY KEY,
            project_id INT NOT NULL,
            project_description TEXT NOT NULL, 
            is_locked BOOLEAN DEFAULT FALSE,
            raw_material_path VARCHAR(255),
            created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (project_id) REFERENCES projects(project_id)
        );

        -- Indexes for the project_details table
        CREATE INDEX idx_project_details_project_id ON project_details(project_id);
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'project_communication') THEN 
        -- Table to store ongoing communication between clients and creators
        CREATE TABLE project_communication (
            communication_id SERIAL PRIMARY KEY,
            project_id INT NOT NULL,
            sender_user_id INT NOT NULL,
            receiver_user_id INT NOT NULL,
            message TEXT NOT NULL,
            communication_timestamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (project_id) REFERENCES projects(project_id),
            FOREIGN KEY (sender_user_id) REFERENCES users(user_id),
            FOREIGN KEY (receiver_user_id) REFERENCES users(user_id)
        );

        -- Indexes for the project_communication table
        CREATE INDEX idx_project_communication_project_id ON project_communication(project_id);
        CREATE INDEX idx_project_communication_sender_user_id ON project_communication(sender_user_id);
        CREATE INDEX idx_project_communication_receiver_user_id ON project_communication(receiver_user_id);
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'project_assignments') THEN 
        -- Table to store project assignments history

        CREATE TYPE assignment_status_enum AS ENUM ('in_progress', 'completed', 'needs_revisions');
        CREATE TABLE project_assignments (
            assignment_id SERIAL PRIMARY KEY,
            project_id INT NOT NULL,
            assignment_order INT,
            old_assigned_creator_id INT,
            new_assigned_creator_id INT,
            assignment_status assignment_status_enum DEFAULT 'in_progress',
            assignment_timestamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (project_id) REFERENCES projects(project_id),
            FOREIGN KEY (old_assigned_creator_id) REFERENCES creators(creator_id),
            FOREIGN KEY (new_assigned_creator_id) REFERENCES creators(creator_id)
        );

        -- Indexes for the project_assignments table
        CREATE INDEX idx_project_assignments_project_id ON project_assignments(project_id);
        CREATE INDEX idx_project_assignments_old_assigned_creator_id ON project_assignments(old_assigned_creator_id);
        CREATE INDEX idx_project_assignments_new_assigned_creator_id ON project_assignments(new_assigned_creator_id);
        CREATE INDEX idx_project_assignments_assignment_status ON project_assignments(assignment_status);
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'task_submissions') THEN 
        -- Table to store task submissions by creators

        CREATE TYPE submission_type_enum AS ENUM ('document', 'code', 'image', 'video');
        CREATE TABLE task_submissions (
            submission_id SERIAL PRIMARY KEY,
            creator_id INT,
            assignment_id INT,
            submission_type submission_type_enum DEFAULT 'document',
            submission_path VARCHAR(255),
            submission_text TEXT NOT NULL,
            submission_timestamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (assignment_id) REFERENCES project_assignments(assignment_id),
            FOREIGN KEY (creator_id) REFERENCES creators(creator_id)
        );

        -- Indexes for the task_submissions table
        CREATE INDEX idx_task_submissions_creator_id ON task_submissions(creator_id);
        CREATE INDEX idx_task_submissions_assignment_id ON task_submissions(assignment_id);
        CREATE INDEX idx_task_submissions_submission_type ON task_submissions(submission_type);
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'client_feedback') THEN 
        -- Table to store feedback from clients
        CREATE TABLE client_feedback (
            feedback_id SERIAL PRIMARY KEY,
            submission_id INT,
            client_id INT,
            feedback_text TEXT NOT NULL,
            feedback_timestamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (submission_id) REFERENCES task_submissions(submission_id),
            FOREIGN KEY (client_id) REFERENCES users(user_id)
        );

        -- Indexes for the client_feedback table
        CREATE INDEX idx_client_feedback_submission_id ON client_feedback(submission_id);
        CREATE INDEX idx_client_feedback_client_id ON client_feedback(client_id);
    END IF;

END $$;