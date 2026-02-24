-- 0001_init.sql
-- Supabase Auth 前提 / RLS enabled
-- Public schema tables:
-- users (profile), companies, applications, interview_types, events, interviews, checklist_items

-- =========
-- Extensions
-- =========
create extension if not exists pgcrypto;

-- =========
-- Helpers: updated_at trigger
-- =========
create or replace function public.set_updated_at()
returns trigger
language plpgsql
as $$
begin
  new.updated_at = now();
  return new;
end;
$$;

-- =========
-- USERS (profile)
-- =========
create table if not exists public.users (
  id uuid primary key references auth.users(id) on delete cascade,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);

drop trigger if exists trg_users_set_updated_at on public.users;
create trigger trg_users_set_updated_at
before update on public.users
for each row execute function public.set_updated_at();

alter table public.users enable row level security;

-- Users policies
drop policy if exists "users_select_own" on public.users;
create policy "users_select_own"
on public.users
for select
using (id = auth.uid());

drop policy if exists "users_insert_own" on public.users;
create policy "users_insert_own"
on public.users
for insert
with check (id = auth.uid());

drop policy if exists "users_update_own" on public.users;
create policy "users_update_own"
on public.users
for update
using (id = auth.uid())
with check (id = auth.uid());

-- =========
-- COMPANIES
-- =========
create table if not exists public.companies (
  id uuid primary key default gen_random_uuid(),
  user_id uuid not null references auth.users(id) on delete cascade,

  name text not null,
  industry text,
  job_type text,
  preference_level int not null default 3 check (preference_level between 1 and 5),
  memo text,

  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now(),
  deleted_at timestamptz
);

create index if not exists companies_user_id_idx on public.companies(user_id);
create index if not exists companies_deleted_at_idx on public.companies(deleted_at);

-- 任意：ユーザーごとの同名企業を防ぐ（soft delete考慮）
-- create unique index if not exists companies_user_name_uniq
-- on public.companies(user_id, name)
-- where deleted_at is null;

drop trigger if exists trg_companies_set_updated_at on public.companies;
create trigger trg_companies_set_updated_at
before update on public.companies
for each row execute function public.set_updated_at();

alter table public.companies enable row level security;

-- Companies policies
drop policy if exists "companies_select_own" on public.companies;
create policy "companies_select_own"
on public.companies
for select
using (user_id = auth.uid());

drop policy if exists "companies_insert_own" on public.companies;
create policy "companies_insert_own"
on public.companies
for insert
with check (user_id = auth.uid());

drop policy if exists "companies_update_own" on public.companies;
create policy "companies_update_own"
on public.companies
for update
using (user_id = auth.uid())
with check (user_id = auth.uid());

drop policy if exists "companies_delete_own" on public.companies;
create policy "companies_delete_own"
on public.companies
for delete
using (user_id = auth.uid());

-- =========
-- APPLICATIONS
-- =========
create table if not exists public.applications (
  id uuid primary key default gen_random_uuid(),
  company_id uuid not null references public.companies(id) on delete cascade,

  application_title text not null,
  status_code text not null default 'APPLIED',
  mypage_url text,
  login_id text,

  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);

create index if not exists applications_company_id_idx on public.applications(company_id);
create index if not exists applications_status_code_idx on public.applications(status_code);

drop trigger if exists trg_applications_set_updated_at on public.applications;
create trigger trg_applications_set_updated_at
before update on public.applications
for each row execute function public.set_updated_at();

alter table public.applications enable row level security;

-- Applications policies
-- NOTE: applications は companies 経由で所有者を判定
drop policy if exists "applications_select_own" on public.applications;
create policy "applications_select_own"
on public.applications
for select
using (
  exists (
    select 1
    from public.companies c
    where c.id = applications.company_id
      and c.user_id = auth.uid()
      and c.deleted_at is null
  )
);

drop policy if exists "applications_insert_own" on public.applications;
create policy "applications_insert_own"
on public.applications
for insert
with check (
  exists (
    select 1
    from public.companies c
    where c.id = applications.company_id
      and c.user_id = auth.uid()
      and c.deleted_at is null
  )
);

drop policy if exists "applications_update_own" on public.applications;
create policy "applications_update_own"
on public.applications
for update
using (
  exists (
    select 1
    from public.companies c
    where c.id = applications.company_id
      and c.user_id = auth.uid()
      and c.deleted_at is null
  )
)
with check (
  exists (
    select 1
    from public.companies c
    where c.id = applications.company_id
      and c.user_id = auth.uid()
      and c.deleted_at is null
  )
);

drop policy if exists "applications_delete_own" on public.applications;
create policy "applications_delete_own"
on public.applications
for delete
using (
  exists (
    select 1
    from public.companies c
    where c.id = applications.company_id
      and c.user_id = auth.uid()
      and c.deleted_at is null
  )
);

-- =========
-- INTERVIEW_TYPES
-- =========
create table if not exists public.interview_types (
  id uuid primary key default gen_random_uuid(),
  user_id uuid not null references auth.users(id) on delete cascade,

  name text not null,
  sort_order int not null default 0,
  is_default boolean not null default false,

  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now(),

  unique(user_id, name)
);

create index if not exists interview_types_user_id_idx on public.interview_types(user_id);

drop trigger if exists trg_interview_types_set_updated_at on public.interview_types;
create trigger trg_interview_types_set_updated_at
before update on public.interview_types
for each row execute function public.set_updated_at();

alter table public.interview_types enable row level security;

-- Interview types policies
drop policy if exists "interview_types_select_own" on public.interview_types;
create policy "interview_types_select_own"
on public.interview_types
for select
using (user_id = auth.uid());

drop policy if exists "interview_types_insert_own" on public.interview_types;
create policy "interview_types_insert_own"
on public.interview_types
for insert
with check (user_id = auth.uid());

drop policy if exists "interview_types_update_own" on public.interview_types;
create policy "interview_types_update_own"
on public.interview_types
for update
using (user_id = auth.uid())
with check (user_id = auth.uid());

drop policy if exists "interview_types_delete_own" on public.interview_types;
create policy "interview_types_delete_own"
on public.interview_types
for delete
using (user_id = auth.uid());

-- =========
-- EVENTS
-- =========
create table if not exists public.events (
  id uuid primary key default gen_random_uuid(),
  application_id uuid not null references public.applications(id) on delete cascade,

  kind text not null, -- INTERVIEW/DEADLINE/TEST/MEETUP...
  title text not null,
  starts_at timestamptz,
  ends_at timestamptz,
  note text,
  external_calendar_event_id text,

  is_done boolean not null default false,
  done_at timestamptz,

  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now(),

  check (ends_at is null or starts_at is null or ends_at >= starts_at),
  check (done_at is null or is_done = true)
);

create index if not exists events_application_id_idx on public.events(application_id);
create index if not exists events_starts_at_idx on public.events(starts_at);
create index if not exists events_is_done_idx on public.events(is_done);
create index if not exists events_kind_idx on public.events(kind);

drop trigger if exists trg_events_set_updated_at on public.events;
create trigger trg_events_set_updated_at
before update on public.events
for each row execute function public.set_updated_at();

alter table public.events enable row level security;

-- Events policies
-- NOTE: events は applications -> companies で所有者判定
drop policy if exists "events_select_own" on public.events;
create policy "events_select_own"
on public.events
for select
using (
  exists (
    select 1
    from public.applications a
    join public.companies c on c.id = a.company_id
    where a.id = events.application_id
      and c.user_id = auth.uid()
      and c.deleted_at is null
  )
);

drop policy if exists "events_insert_own" on public.events;
create policy "events_insert_own"
on public.events
for insert
with check (
  exists (
    select 1
    from public.applications a
    join public.companies c on c.id = a.company_id
    where a.id = events.application_id
      and c.user_id = auth.uid()
      and c.deleted_at is null
  )
);

drop policy if exists "events_update_own" on public.events;
create policy "events_update_own"
on public.events
for update
using (
  exists (
    select 1
    from public.applications a
    join public.companies c on c.id = a.company_id
    where a.id = events.application_id
      and c.user_id = auth.uid()
      and c.deleted_at is null
  )
)
with check (
  exists (
    select 1
    from public.applications a
    join public.companies c on c.id = a.company_id
    where a.id = events.application_id
      and c.user_id = auth.uid()
      and c.deleted_at is null
  )
);

drop policy if exists "events_delete_own" on public.events;
create policy "events_delete_own"
on public.events
for delete
using (
  exists (
    select 1
    from public.applications a
    join public.companies c on c.id = a.company_id
    where a.id = events.application_id
      and c.user_id = auth.uid()
      and c.deleted_at is null
  )
);

-- =========
-- INTERVIEWS (extends EVENTS)
-- =========
create table if not exists public.interviews (
  id uuid primary key default gen_random_uuid(),
  event_id uuid not null unique references public.events(id) on delete cascade,

  round int not null check (round >= 1),
  interview_type_id uuid not null references public.interview_types(id),

  memo text,

  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);

create index if not exists interviews_interview_type_id_idx on public.interviews(interview_type_id);

drop trigger if exists trg_interviews_set_updated_at on public.interviews;
create trigger trg_interviews_set_updated_at
before update on public.interviews
for each row execute function public.set_updated_at();

alter table public.interviews enable row level security;

-- Interviews policies
-- NOTE: interviews は events -> applications -> companies で所有者判定
drop policy if exists "interviews_select_own" on public.interviews;
create policy "interviews_select_own"
on public.interviews
for select
using (
  exists (
    select 1
    from public.events e
    join public.applications a on a.id = e.application_id
    join public.companies c on c.id = a.company_id
    where e.id = interviews.event_id
      and c.user_id = auth.uid()
      and c.deleted_at is null
  )
);

drop policy if exists "interviews_insert_own" on public.interviews;
create policy "interviews_insert_own"
on public.interviews
for insert
with check (
  exists (
    select 1
    from public.events e
    join public.applications a on a.id = e.application_id
    join public.companies c on c.id = a.company_id
    where e.id = interviews.event_id
      and c.user_id = auth.uid()
      and c.deleted_at is null
  )
);

drop policy if exists "interviews_update_own" on public.interviews;
create policy "interviews_update_own"
on public.interviews
for update
using (
  exists (
    select 1
    from public.events e
    join public.applications a on a.id = e.application_id
    join public.companies c on c.id = a.company_id
    where e.id = interviews.event_id
      and c.user_id = auth.uid()
      and c.deleted_at is null
  )
)
with check (
  exists (
    select 1
    from public.events e
    join public.applications a on a.id = e.application_id
    join public.companies c on c.id = a.company_id
    where e.id = interviews.event_id
      and c.user_id = auth.uid()
      and c.deleted_at is null
  )
);

drop policy if exists "interviews_delete_own" on public.interviews;
create policy "interviews_delete_own"
on public.interviews
for delete
using (
  exists (
    select 1
    from public.events e
    join public.applications a on a.id = e.application_id
    join public.companies c on c.id = a.company_id
    where e.id = interviews.event_id
      and c.user_id = auth.uid()
      and c.deleted_at is null
  )
);

-- =========
-- CHECKLIST_ITEMS
-- =========
create table if not exists public.checklist_items (
  id uuid primary key default gen_random_uuid(),
  application_id uuid not null references public.applications(id) on delete cascade,

  title text not null,
  is_done boolean not null default false,
  sort_order int not null default 0,

  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);

create index if not exists checklist_items_application_id_idx on public.checklist_items(application_id);
create index if not exists checklist_items_is_done_idx on public.checklist_items(is_done);

drop trigger if exists trg_checklist_items_set_updated_at on public.checklist_items;
create trigger trg_checklist_items_set_updated_at
before update on public.checklist_items
for each row execute function public.set_updated_at();

alter table public.checklist_items enable row level security;

-- Checklist policies
drop policy if exists "checklist_select_own" on public.checklist_items;
create policy "checklist_select_own"
on public.checklist_items
for select
using (
  exists (
    select 1
    from public.applications a
    join public.companies c on c.id = a.company_id
    where a.id = checklist_items.application_id
      and c.user_id = auth.uid()
      and c.deleted_at is null
  )
);

drop policy if exists "checklist_insert_own" on public.checklist_items;
create policy "checklist_insert_own"
on public.checklist_items
for insert
with check (
  exists (
    select 1
    from public.applications a
    join public.companies c on c.id = a.company_id
    where a.id = checklist_items.application_id
      and c.user_id = auth.uid()
      and c.deleted_at is null
  )
);

drop policy if exists "checklist_update_own" on public.checklist_items;
create policy "checklist_update_own"
on public.checklist_items
for update
using (
  exists (
    select 1
    from public.applications a
    join public.companies c on c.id = a.company_id
    where a.id = checklist_items.application_id
      and c.user_id = auth.uid()
      and c.deleted_at is null
  )
)
with check (
  exists (
    select 1
    from public.applications a
    join public.companies c on c.id = a.company_id
    where a.id = checklist_items.application_id
      and c.user_id = auth.uid()
      and c.deleted_at is null
  )
);

drop policy if exists "checklist_delete_own" on public.checklist_items;
create policy "checklist_delete_own"
on public.checklist_items
for delete
using (
  exists (
    select 1
    from public.applications a
    join public.companies c on c.id = a.company_id
    where a.id = checklist_items.application_id
      and c.user_id = auth.uid()
      and c.deleted_at is null
  )
);