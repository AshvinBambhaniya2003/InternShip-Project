-- Make sure to attach design image/pdf in the same folder.
-- Write your DDL queries here.

// Use DBML to define your database structure
// Docs: https://dbml.dbdiagram.io/docs

Table users {
  user_id integer [primary key]
  username varchar [unique, not null]
  role ENUM('client', 'creator') [not null]
  name varchar [not null]
}


Table Skills {
  skill_id integer [primary key]
  skill varchar [not null]
  description VARCHAR
}

Table Demos {
  demo_id integer [primary key]
  creator_id integer
  skill_id integer
  demo_type varchar
  demo_path varchar
}

Table Categories {
  category_id integer [primary key]
  category_name varchar [not null,unique]
}

TABLE creator_Skills {
    creator_id integer [not null]
    skill_id integer [not null]
    Indexes {
    (creator_id, skill_id) [pk] 
  }
}

Ref: creator_Skills.(creator_id, skill_id) - Demos.(creator_id, skill_id)

Table Projects {
  project_id integer [primary key]
  client_id integer [not null]
  project_name varchar [not null]
  project_description varchar [not null]
  project_status ENUM('open', 'locked') [not null] 
}

Table ProjectMaterials {
  material_id integer [primary key]
  project_id integer [not null]
  material_type varchar
  path varchar
}

TABLE Assignments {
  assignment_id integer [primary key]
  project_id integer [not null]
  creator_id integer [not null]
  assignee_id INT
  is_locked BOOLEAN [not null]
  deadline date
  submission_path VARCHAR 
  is_done BOOLEAN [not null]
}


Table AssignmentMaterials {
  assigment_material_id integer [primary key]
  assignment_id integer [not null]
  material_type varchar
  path varchar
}

Table CommunicationMessages {
  message_id INT [primary key]
  communication_id INT [not null, unique]
  sender_id INT [not null, ref: > users.user_id]
  receiver_id INT [not null, ref: > users.user_id]
  message TEXT [not null]
  sent_at TIMESTAMP
  is_read BOOLEAN 
}


Table Communications {
  communication_id INT [PRIMARY KEY]
  assignment_id INT [NOT NULL]
}

Ref: "CommunicationMessages"."communication_id" - "Communications"."communication_id"

Ref: "users"."user_id" < "creator_Skills"."creator_id"

Ref: "users"."user_id" < "Projects"."client_id"

Ref: "Projects"."project_id" < "ProjectMaterials"."project_id"

Ref: "Categories"."category_name" < "Skills"."skill"

Ref: "Assignments"."assignment_id" < "AssignmentMaterials"."assignment_id"

Ref: "Projects"."project_id" < "Assignments"."project_id"

Ref: "Skills"."skill_id" < "creator_Skills"."skill_id"

Ref: "Assignments"."assignment_id" < "Communications"."assignment_id"

Ref: "users"."user_id" < "Assignments"."creator_id"