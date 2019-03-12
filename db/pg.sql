CREATE TABLE project (
  project_id serial NOT NULL PRIMARY KEY,
  project_uuid varchar(50) NOT NULL UNIQUE,
  project_name varchar(200) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  last_actioned_at timestamp,
  project_status varchar(50) NOT NULL,
	CONSTRAINT valid_dates CHECK (last_actioned_at > created_at)
);

/* Add user to the workspace*/
CREATE TABLE workspace (
  project_uuid varchar(50) NOT NULL,
  feature_uuid varchar(50),
  wireframe_uuid varchar(50),
  active_tab varchar(50) NOT NULL
);

CREATE TABLE feature (
  feature_id serial NOT NULL PRIMARY KEY,
  project_uuid varchar(50) NOT NULL REFERENCES project(project_uuid),
  feature_uuid varchar(50) NOT NULL UNIQUE,
  feature_title text NOT NULL,
  feature_content text,
  created_at timestamp NOT NULL DEFAULT NOW(),
  last_actioned_at timestamp,
  deleted boolean DEFAULT FALSE,
	CONSTRAINT valid_dates CHECK (last_actioned_at > created_at)
);

CREATE TABLE wireframe (
  wireframe_id serial NOT NULL PRIMARY KEY,
  project_uuid varchar(50) NOT NULL,
  feature_uuid varchar(50) NOT NULL REFERENCES feature(feature_uuid),
  wireframe_uuid varchar(50) NOT NULL UNIQUE,
  wireframe_title text NOT NULL,
  wireframe_content text,
  created_at timestamp NOT NULL DEFAULT NOW(),
  last_actioned_at timestamp,
  deleted boolean DEFAULT FALSE,
	CONSTRAINT valid_dates CHECK (last_actioned_at > created_at)
);

CREATE TABLE Element (
  wireframe_uuid varchar(50) NOT NULL REFERENCES wireframe(wireframe_uuid),
  timestampkey    int NOT NULL,
	element_type    varchar(50) NOT NULL,
	stroke_style    varchar(50),
	fill_style      varchar(50),
	start_x         int NOT NULL,
	start_y         int NOT NULL,
	width           int,
	height          int,
	element_text    varchar(50) NOT NULL,
	checked         boolean,
	min_width       int,
	min_height      int,
	rounded_corner   boolean,
	font_size        int,
	foreground_color varchar(50),
	resizable       boolean,
	border          varchar(50),
	editable        boolean
);
