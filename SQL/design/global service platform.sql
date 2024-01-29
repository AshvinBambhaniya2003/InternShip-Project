-- Make sure to attach design image/pdf in the same folder.
-- Write your DDL queries here.

// Use DBML to define your database structure
// Docs: https://dbml.dbdiagram.io/docs

Table users {
  user_id integer [primary key]
  username varchar
  role varchar 
  name varchar
}

Table client {
  client_id integer [primary key]
  username varchar
}

Table creators {
  creator_id integer [primary key]
  username varchar
}



Table Skills {
  skill_id integer [primary key]
  creator_id integer 
  skill varchar
}

Table Demos {
  demo_id integer [primary key]
  creator_id integer
  demo_type varchar
  demo_path varchar
}

Table Categories {
  category_id integer [primary key]
  category_name varchar
}

TABLE creator_Skills {
    creator_id integer
    skill_id integer
}

Table Projects {
  project_id integer [primary key]
  client_id integer 
  project_name varchar
  project_description varchar
}

Table ProjectMaterials {
  material_id integer [primary key]
  project_id integer 
  material_type varchar
  path varchar
}

TABLE Assignments {
  assignment_id integer [primary key]
  project_id integer 
  creator_id integer 
  status varchar
}

TABLE Communications {
  communication_id INT [PRIMARY KEY]
  assignment_id INT
  message TEXT
}

TABLE Feedback {
  feedback_id INT [PRIMARY KEY]
  assignment_id INT
  client_feedback TEXT
}

TABLE Deadlines {
  deadline_id INT [PRIMARY KEY]
  assignment_id INT
}



Ref: "users"."user_id" < "client"."client_id"

Ref: "users"."user_id" < "creators"."creator_id"

Ref: "client"."client_id" < "Projects"."client_id"

Ref: "creators"."creator_id" < "Demos"."creator_id"

Ref: "Projects"."project_id" < "Assignments"."project_id"

Ref: "creators"."creator_id" < "Assignments"."creator_id"

Ref: "Categories"."category_name" < "Skills"."skill"

Ref: "creators"."creator_id" < "Skills"."creator_id"

Ref: "creators"."creator_id" < "creator_Skills"."creator_id"

Ref: "Skills"."skill_id" < "creator_Skills"."skill_id"

Ref: "Assignments"."assignment_id" < "Communications"."assignment_id"

Ref: "Assignments"."assignment_id" < "Feedback"."assignment_id"

Ref: "Assignments"."assignment_id" < "Deadlines"."assignment_id"

Ref: "Projects"."project_id" < "ProjectMaterials"."project_id"