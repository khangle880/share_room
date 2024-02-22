-- +goose Up
CREATE TYPE user_role AS ENUM ('admin', 'user');
CREATE TYPE room_role AS ENUM ('admin', 'member');
CREATE TYPE task_role AS ENUM ('assignee', 'reviewer', 'creator');
CREATE TYPE trans_role AS ENUM ('creator', 'partner');
CREATE TYPE budget_role AS ENUM ('owner', 'contributor');
CREATE TYPE period_type AS ENUM ('weekly', 'monthly', 'yearly');
CREATE TYPE status_type AS ENUM ('pending', 'in_progress', 'completed');
CREATE TYPE category_type AS ENUM ('expense', 'income');

-- +goose Down
DROP TYPE user_role;
DROP TYPE user_role;
DROP TYPE room_role;
DROP TYPE task_role;
DROP TYPE trans_role;
DROP TYPE budget_role;
DROP TYPE period_type;
DROP TYPE status_type;
DROP TYPE category_type;