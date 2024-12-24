package postgres

const (
	queryInitUsers = `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		role TEXT NOT NULL DEFAULT 'customer',
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	
	CREATE INDEX idx_users_username ON users(username);
	CREATE INDEX idx_users_email ON users(email);
	`

	queryInitTattoos = `
	CREATE TABLE IF NOT EXISTS tattoos (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT,
		price DECIMAL(10, 2) NOT NULL,
		image_url TEXT,
		artist_id INTEGER REFERENCES users(id),
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE INDEX idx_tattoos_artist_id ON tattoos(artist_id);
	`

	queryInitCart = `
	CREATE TABLE IF NOT EXISTS cart (
		id SERIAL PRIMARY KEY,
		user_id INTEGER NOT NULL REFERENCES users(id),
		tattoo_id INTEGER NOT NULL REFERENCES tattoos(id),
		quantity INTEGER NOT NULL DEFAULT 1,
		added_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		UNIQUE (user_id, tattoo_id)
	);

	CREATE INDEX idx_cart_user_id ON cart(user_id);
	CREATE INDEX idx_cart_tattoo_id ON cart(tattoo_id);
	`

	queryInitOrders = `
	CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		customer_id INTEGER NOT NULL REFERENCES users(id),
		order_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		status TEXT NOT NULL DEFAULT 'pending',
		total_cost DECIMAL(10, 2) NOT NULL,
		completed_at TIMESTAMP
	);

	CREATE INDEX idx_orders_customer_id ON orders(customer_id);
	CREATE INDEX idx_orders_status ON orders(status);
	`

	queryInitOrderItems = `
	CREATE TABLE IF NOT EXISTS order_items (
		id SERIAL PRIMARY KEY,
		order_id INTEGER NOT NULL REFERENCES orders(id),
		tattoo_id INTEGER NOT NULL REFERENCES tattoos(id),
		quantity INTEGER NOT NULL DEFAULT 1,
		item_price DECIMAL(10, 2) NOT NULL,
		UNIQUE (order_id, tattoo_id)
	);

	CREATE INDEX idx_order_items_order_id ON order_items(order_id);
	CREATE INDEX idx_order_items_tattoo_id ON order_items(tattoo_id);
	`

	queryInitTags = `
	CREATE TABLE IF NOT EXISTS tags (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL UNIQUE
	);

	CREATE INDEX idx_tags_name ON tags(name);
	`

	queryInitTattooTags = `
	CREATE TABLE IF NOT EXISTS tattoo_tags (
		tattoo_id INTEGER NOT NULL REFERENCES tattoos(id),
		tag_id INTEGER NOT NULL REFERENCES tags(id),
		PRIMARY KEY (tattoo_id, tag_id)
	);

	CREATE INDEX idx_tattoo_tags_tattoo_id ON tattoo_tags(tattoo_id);
	CREATE INDEX idx_tattoo_tags_tag_id ON tattoo_tags(tag_id);
	`
)
