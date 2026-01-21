CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    role VARCHAR(50) NOT NULL DEFAULT 'user',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT check_role CHECK (role IN ('user', 'operator', 'admin'))
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role ON users(role);

CREATE TABLE notification_channels (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type VARCHAR(50) NOT NULL,
    target VARCHAR(255) NOT NULL,
    is_enabled BOOLEAN NOT NULL DEFAULT true,
    is_verified BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT check_channel_type CHECK (type IN ('email', 'telegram'))
);

CREATE INDEX idx_notification_channels_user_id ON notification_channels(user_id);
CREATE INDEX idx_notification_channels_type ON notification_channels(type);
CREATE INDEX idx_notification_channels_is_enabled ON notification_channels(is_enabled);

CREATE TABLE subscriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_subscriptions_user_id ON subscriptions(user_id);

CREATE TABLE service_groups (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    "order" INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_service_groups_slug ON service_groups(slug);
CREATE INDEX idx_service_groups_order ON service_groups("order");

CREATE TABLE services (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    status VARCHAR(50) NOT NULL DEFAULT 'operational',
    group_id UUID REFERENCES service_groups(id) ON DELETE SET NULL,
    "order" INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT check_status CHECK (status IN ('operational', 'degraded', 'partial_outage', 'major_outage', 'maintenance'))
);

CREATE INDEX idx_services_slug ON services(slug);
CREATE INDEX idx_services_group_id ON services(group_id);
CREATE INDEX idx_services_status ON services(status);
CREATE INDEX idx_services_order ON services("order");

CREATE TABLE service_tags (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    service_id UUID NOT NULL REFERENCES services(id) ON DELETE CASCADE,
    key VARCHAR(100) NOT NULL,
    value TEXT NOT NULL,
    CONSTRAINT unique_service_tag UNIQUE (service_id, key)
);

CREATE INDEX idx_service_tags_service_id ON service_tags(service_id);
CREATE INDEX idx_service_tags_key ON service_tags(key);

CREATE TABLE subscription_services (
    subscription_id UUID NOT NULL REFERENCES subscriptions(id) ON DELETE CASCADE,
    service_id UUID NOT NULL REFERENCES services(id) ON DELETE CASCADE,
    PRIMARY KEY (subscription_id, service_id)
);

CREATE INDEX idx_subscription_services_service_id ON subscription_services(service_id);

CREATE TABLE event_templates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slug VARCHAR(255) NOT NULL UNIQUE,
    type VARCHAR(50) NOT NULL,
    title_template TEXT NOT NULL,
    body_template TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT check_template_type CHECK (type IN ('incident', 'maintenance'))
);

CREATE INDEX idx_event_templates_slug ON event_templates(slug);
CREATE INDEX idx_event_templates_type ON event_templates(type);

CREATE TABLE events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(500) NOT NULL,
    type VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    severity VARCHAR(50),
    description TEXT,
    started_at TIMESTAMP,
    resolved_at TIMESTAMP,
    scheduled_start_at TIMESTAMP,
    scheduled_end_at TIMESTAMP,
    notify_subscribers BOOLEAN NOT NULL DEFAULT false,
    template_id UUID REFERENCES event_templates(id) ON DELETE SET NULL,
    created_by UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT check_event_type CHECK (type IN ('incident', 'maintenance')),
    CONSTRAINT check_severity CHECK (severity IS NULL OR severity IN ('minor', 'major', 'critical'))
);

CREATE INDEX idx_events_type ON events(type);
CREATE INDEX idx_events_status ON events(status);
CREATE INDEX idx_events_severity ON events(severity);
CREATE INDEX idx_events_created_by ON events(created_by);
CREATE INDEX idx_events_created_at ON events(created_at DESC);
CREATE INDEX idx_events_started_at ON events(started_at DESC);
CREATE INDEX idx_events_scheduled_start_at ON events(scheduled_start_at);

CREATE TABLE event_services (
    event_id UUID NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    service_id UUID NOT NULL REFERENCES services(id) ON DELETE CASCADE,
    PRIMARY KEY (event_id, service_id)
);

CREATE INDEX idx_event_services_service_id ON event_services(service_id);

CREATE TABLE event_updates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id UUID NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    status VARCHAR(50) NOT NULL,
    message TEXT NOT NULL,
    notify_subscribers BOOLEAN NOT NULL DEFAULT false,
    created_by UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_event_updates_event_id ON event_updates(event_id);
CREATE INDEX idx_event_updates_created_at ON event_updates(created_at DESC);
