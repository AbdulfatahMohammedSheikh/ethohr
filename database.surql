
-- simple configuration and database setup

-------------
-- Roles table
-------------
BEGIN TRANSACTION;

define table roles SCHEMAFULL;

define  field name on table roles type string;
define  field isband on table roles type bool;

DEFINE INDEX roleName ON TABLE roles COLUMNS name UNIQUE;

COMMIT TRANSACTION;

-- TODO: add events for deleting roles


--------------------------------------------------------

-------------
-- Tag table
-------------
BEGIN TRANSACTION;

define table tags SCHEMAFULL;

define  field name on table tags type string;

DEFINE INDEX tagName ON TABLE tags  COLUMNS name UNIQUE;

COMMIT TRANSACTION;

-- Events 

-- DEFINE EVENT tag_deleted ON TABLE tags WHEN $event = "DELETE" THEN (
-- update users  set tags -= $before.id ;
-- );

DEFINE EVENT user_updated ON TABLE user WHEN $event = "UPDATE" THEN (
    update  users  SET tags -= $this.id
);
--------------------------------------------------------
