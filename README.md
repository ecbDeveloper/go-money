<div id="top">

<!-- HEADER STYLE: CLASSIC -->
<div align="center">

<img src="readmeai/assets/logos/purple.svg" width="30%" style="position: relative; top: 0; right: 0;" alt="Project Logo"/>

# GO-MONEY

<em></em>

<!-- BADGES -->
<img src="https://img.shields.io/github/license/ecbDeveloper/go-money?style=default&logo=opensourceinitiative&logoColor=white&color=0080ff" alt="license">
<img src="https://img.shields.io/github/last-commit/ecbDeveloper/go-money?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">
<img src="https://img.shields.io/github/languages/top/ecbDeveloper/go-money?style=default&color=0080ff" alt="repo-top-language">
<img src="https://img.shields.io/github/languages/count/ecbDeveloper/go-money?style=default&color=0080ff" alt="repo-language-count">

<!-- default option, no dependency badges. -->

<!-- default option, no dependency badges. -->

</div>
<br>

## Overview
A financial transactions API built in Golang, enabling secure and efficient management of transfers, deposits, and withdrawals, with robust concurrency control and real-time processing to ensure accurate and reliable operations.

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Overview](#overview)
- [Project Structure](#project-structure)
    - [Project Index](#project-index)
- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [Usage](#usage)

---

## Project Structure

```sh
└── go-money/
    ├── cmd
    │   └── gomoney
    ├── compose.yml
    ├── go.mod
    ├── go.sum
    └── internal
        ├── api
        ├── db
        ├── models
        ├── services
        └── shared
```

### Project Index

<details open>
	<summary><b><code>GO-MONEY/</code></b></summary>
	<!-- __root__ Submodule -->
	<details>
		<summary><b>__root__</b></summary>
		<blockquote>
			<div class='directory-path' style='padding: 8px 0; color: #666;'>
				<code><b>⦿ __root__</b></code>
			<table style='width: 100%; border-collapse: collapse;'>
			<thead>
				<tr style='background-color: #f8f9fa;'>
					<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
					<th style='text-align: left; padding: 8px;'>Summary</th>
				</tr>
			</thead>
				<tr style='border-bottom: 1px solid #eee;'>
					<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/go.mod'>go.mod</a></b></td>
					<td style='padding: 8px;'>- The <code>go.mod</code> file defines the projects module path and specifies its dependencies<br>- It declares the project as <code>github.com/ecbDeveloper/go-money</code> and lists required packages for session management, routing, UUID generation, CSRF protection, PostgreSQL database interaction, environment variable loading, and cryptography<br>- These dependencies support the applications core functionality.</td>
				</tr>
				<tr style='border-bottom: 1px solid #eee;'>
					<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/go.sum'>go.sum</a></b></td>
					<td style='padding: 8px;'>- The <code>go.sum</code> file records the cryptographic checksums of the <code>pgxstore</code> package from the <code>alexedwards/scs</code> library<br>- This ensures that the correct, unaltered version of this session store (used for managing user sessions, likely within a larger web application) is used, maintaining the integrity and security of the application<br>- Its a crucial part of the projects dependency management system, preventing dependency vulnerabilities.</td>
				</tr>
				<tr style='border-bottom: 1px solid #eee;'>
					<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/compose.yml'>compose.yml</a></b></td>
					<td style='padding: 8px;'>- Compose defines a PostgreSQL database service named db<br>- It uses the latest PostgreSQL image, maps port 5431 to 5432, and sets environment variables for user, password, and database name (go-money)<br>- A persistent volume db ensures data preservation across container restarts, contributing to the application's persistent data storage within the overall project architecture.</td>
				</tr>
			</table>
		</blockquote>
	</details>
	<!-- cmd Submodule -->
	<details>
		<summary><b>cmd</b></summary>
		<blockquote>
			<div class='directory-path' style='padding: 8px 0; color: #666;'>
				<code><b>⦿ cmd</b></code>
			<!-- gomoney Submodule -->
			<details>
				<summary><b>gomoney</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ cmd.gomoney</b></code>
					<table style='width: 100%; border-collapse: collapse;'>
					<thead>
						<tr style='background-color: #f8f9fa;'>
							<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
							<th style='text-align: left; padding: 8px;'>Summary</th>
						</tr>
					</thead>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/cmd/gomoney/main.go'>main.go</a></b></td>
							<td style='padding: 8px;'>- Gomoney`s main function initializes a web server<br>- It establishes a database connection, configures session management using PostgreSQL, and initializes API services for clients and accounts<br>- The server then binds API routes and listens for incoming requests on port 8082, acting as the applications entry point<br>- Environment variables manage database credentials.</td>
						</tr>
					</table>
				</blockquote>
			</details>
		</blockquote>
	</details>
	<!-- internal Submodule -->
	<details>
		<summary><b>internal</b></summary>
		<blockquote>
			<div class='directory-path' style='padding: 8px 0; color: #666;'>
				<code><b>⦿ internal</b></code>
			<!-- shared Submodule -->
			<details>
				<summary><b>shared</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ internal.shared</b></code>
					<table style='width: 100%; border-collapse: collapse;'>
					<thead>
						<tr style='background-color: #f8f9fa;'>
							<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
							<th style='text-align: left; padding: 8px;'>Summary</th>
						</tr>
					</thead>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/shared/validators.go'>validators.go</a></b></td>
							<td style='padding: 8px;'>- Validators.go provides reusable input validation functions for the application<br>- It offers checks for blank strings, valid email addresses, and string length constraints (minimum and maximum character counts)<br>- These functions are centrally located within the <code>internal/shared</code> package, promoting code reusability and maintainability across the entire project<br>- This ensures consistent data validation throughout the applications various components.</td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/shared/converters.go'>converters.go</a></b></td>
							<td style='padding: 8px;'>- Converters.go<code> provides data type conversion functions within the </code>internal/shared<code> package<br>- It facilitates interoperability between PostgreSQLs </code>numeric<code> type and Gos </code>float64` type, enabling seamless data exchange between the database and application layers<br>- These conversions are crucial for consistent data handling throughout the application.</td>
						</tr>
					</table>
				</blockquote>
			</details>
			<!-- services Submodule -->
			<details>
				<summary><b>services</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ internal.services</b></code>
					<table style='width: 100%; border-collapse: collapse;'>
					<thead>
						<tr style='background-color: #f8f9fa;'>
							<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
							<th style='text-align: left; padding: 8px;'>Summary</th>
						</tr>
					</thead>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/services/conta_services.go'>conta_services.go</a></b></td>
							<td style='padding: 8px;'>- Conta_services.go<code> provides account management functionalities within the </code>go-money<code> application<br>- It offers services for account creation, balance retrieval, transactions (deposits, withdrawals, transfers), and deletion<br>- The service interacts with a database using </code>sqlc` for data persistence and transaction management, ensuring data integrity<br>- Error handling is implemented to manage various scenarios, such as insufficient funds or invalid operations.</td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/services/cliente_services.go'>cliente_services.go</a></b></td>
							<td style='padding: 8px;'>- Cliente_services.go<code> provides client management functionalities within the </code>go-money<code> application<br>- It handles client creation, securely storing passwords using bcrypt, and authenticating clients based on email and password credentials<br>- The service interacts with a PostgreSQL database via </code>sqlc` for data persistence, managing different client types (individuals and businesses) and enforcing data uniqueness constraints.</td>
						</tr>
					</table>
				</blockquote>
			</details>
			<!-- api Submodule -->
			<details>
				<summary><b>api</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ internal.api</b></code>
					<table style='width: 100%; border-collapse: collapse;'>
					<thead>
						<tr style='background-color: #f8f9fa;'>
							<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
							<th style='text-align: left; padding: 8px;'>Summary</th>
						</tr>
					</thead>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/api/conta_handler.go'>conta_handler.go</a></b></td>
							<td style='padding: 8px;'>- Conta_handler.go<code> provides HTTP handlers for account management within the </code>go-money<code> API<br>- It exposes endpoints to create, retrieve balance, perform transactions (deposits, withdrawals, transfers), and delete accounts<br>- Each handler interacts with the </code>AccountService` to manage account data and responds with appropriate HTTP status codes and JSON-formatted messages, indicating success or specific error conditions.</td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/api/api.go'>api.go</a></b></td>
							<td style='padding: 8px;'>- Api.go` defines the API struct, serving as the central component for the applications API layer<br>- It aggregates core services (ClientService, AccountService) and session management (SessionManager), using Chi router for request handling<br>- The struct facilitates interaction between the applications business logic and external clients.</td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/api/auth.go'>auth.go</a></b></td>
							<td style='padding: 8px;'>- Auth.go<code> provides authentication middleware and CSRF token handling for the </code>internal/api` package<br>- The middleware verifies user sessions, rejecting unauthenticated requests<br>- A dedicated handler generates and returns CSRF tokens, crucial for protecting against cross-site request forgery attacks within the web applications API<br>- This ensures secure access control throughout the application.</td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/api/cliente_handler.go'>cliente_handler.go</a></b></td>
							<td style='padding: 8px;'>- Cliente_handler.go<code> implements HTTP handlers for client-related API endpoints within the </code>go-money<code> application<br>- It manages client creation, validating input data, and handling potential errors<br>- The handlers interact with the </code>services<code> layer for business logic and utilize the </code>models` layer for data structures, providing JSON responses to client requests for account creation and login<br>- Session management is also integrated for authentication.</td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/api/routes.go'>routes.go</a></b></td>
							<td style='padding: 8px;'>- Routes.go defines the API routing for a Go web application<br>- It configures middleware for request logging, error handling, session management, and CSRF protection<br>- The code establishes API endpoints for client creation, login, account management (creation, retrieval, deletion, and transactions), and balance inquiries<br>- Authentication middleware protects account-related routes<br>- The API is versioned (v1).</td>
						</tr>
					</table>
				</blockquote>
			</details>
			<!-- models Submodule -->
			<details>
				<summary><b>models</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ internal.models</b></code>
					<table style='width: 100%; border-collapse: collapse;'>
					<thead>
						<tr style='background-color: #f8f9fa;'>
							<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
							<th style='text-align: left; padding: 8px;'>Summary</th>
						</tr>
					</thead>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/models/cliente.go'>cliente.go</a></b></td>
							<td style='padding: 8px;'>- Cliente.go<code> defines data structures for client creation and authentication within the </code>internal/models<code> package<br>- It provides </code>CreateClient<code> and </code>AuthenticateClient` structs, representing client registration and login requests respectively<br>- Crucially, it includes validation functions for each struct, ensuring data integrity before processing, thus contributing to robust data handling within the application.</td>
						</tr>
						<tr style='border-bottom: 1px solid #eee;'>
							<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/models/conta.go'>conta.go</a></b></td>
							<td style='padding: 8px;'>- Conta.go<code> defines data structures for account transactions and money transfers within the </code>models` package<br>- It provides validation functions for these structures, ensuring data integrity before processing<br>- These models likely serve as input for other parts of the application responsible for handling financial operations, contributing to the overall systems data management and validation layer.</td>
						</tr>
					</table>
				</blockquote>
			</details>
			<!-- db Submodule -->
			<details>
				<summary><b>db</b></summary>
				<blockquote>
					<div class='directory-path' style='padding: 8px 0; color: #666;'>
						<code><b>⦿ internal.db</b></code>
					<!-- queries Submodule -->
					<details>
						<summary><b>queries</b></summary>
						<blockquote>
							<div class='directory-path' style='padding: 8px 0; color: #666;'>
								<code><b>⦿ internal.db.queries</b></code>
							<table style='width: 100%; border-collapse: collapse;'>
							<thead>
								<tr style='background-color: #f8f9fa;'>
									<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
									<th style='text-align: left; padding: 8px;'>Summary</th>
								</tr>
							</thead>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/queries/cliente.sql'>cliente.sql</a></b></td>
									<td style='padding: 8px;'>- The <code>cliente.sql</code> file provides SQL queries for interacting with the <code>cliente</code> table within the database<br>- It facilitates client creation via <code>CreateClient</code> and retrieval by email using <code>GetClientByEmail</code><br>- These functions are crucial for user account management within the broader application architecture.</td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/queries/pessoa_fisica.sql'>pessoa_fisica.sql</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/queries/pessoa_juridica.sql'>pessoa_juridica.sql</a></b></td>
									<td style='padding: 8px;'>- The <code>pessoa_juridica.sql</code> file provides a SQL query for creating new legal entity records in the database<br>- Its part of the <code>internal/db/queries</code> directory, suggesting a role within the broader database interaction layer of the application<br>- The query inserts data into the <code>pessoa_juridica</code> table, populating fields such as client ID, creation date, trade name, and CNPJ<br>- This function supports the core applications ability to manage legal entity information.</td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/queries/transferencia.sql'>transferencia.sql</a></b></td>
									<td style='padding: 8px;'>- The <code>transferencia.sql</code> file defines a SQL query for creating new transfer records within the database<br>- Its part of the <code>internal/db/queries</code> module, contributing to the projects data persistence layer<br>- This query inserts data into the <code>transferencia</code> table, specifying account ID, transaction value, and type<br>- The function facilitates the core banking operation of transferring funds.</td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/queries/conta.sql'>conta.sql</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
							</table>
						</blockquote>
					</details>
					<!-- sqlc Submodule -->
					<details>
						<summary><b>sqlc</b></summary>
						<blockquote>
							<div class='directory-path' style='padding: 8px 0; color: #666;'>
								<code><b>⦿ internal.db.sqlc</b></code>
							<table style='width: 100%; border-collapse: collapse;'>
							<thead>
								<tr style='background-color: #f8f9fa;'>
									<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
									<th style='text-align: left; padding: 8px;'>Summary</th>
								</tr>
							</thead>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/sqlc/pessoa_juridica.sql.go'>pessoa_juridica.sql.go</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/sqlc/transferencia.sql.go'>transferencia.sql.go</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/sqlc/sqlc.yaml'>sqlc.yaml</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/sqlc/db.go'>db.go</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/sqlc/cliente.sql.go'>cliente.sql.go</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/sqlc/conta.sql.go'>conta.sql.go</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/sqlc/pessoa_fisica.sql.go'>pessoa_fisica.sql.go</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/sqlc/models.go'>models.go</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
							</table>
						</blockquote>
					</details>
					<!-- migrations Submodule -->
					<details>
						<summary><b>migrations</b></summary>
						<blockquote>
							<div class='directory-path' style='padding: 8px 0; color: #666;'>
								<code><b>⦿ internal.db.migrations</b></code>
							<table style='width: 100%; border-collapse: collapse;'>
							<thead>
								<tr style='background-color: #f8f9fa;'>
									<th style='width: 30%; text-align: left; padding: 8px;'>File Name</th>
									<th style='text-align: left; padding: 8px;'>Summary</th>
								</tr>
							</thead>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/migrations/007_create_transferencias.sql'>007_create_transferencias.sql</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/migrations/009_create_status_conta.sql'>009_create_status_conta.sql</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/migrations/010_add_column_status_on_table_conta.sql'>010_add_column_status_on_table_conta.sql</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/migrations/004_create_pessoa_fisica.sql'>004_create_pessoa_fisica.sql</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/migrations/008_create_sessions.sql'>008_create_sessions.sql</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/migrations/001_create_categoria_usuario.sql'>001_create_categoria_usuario.sql</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/migrations/002_create_cliente.sql'>002_create_cliente.sql</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/migrations/003_create_conta.sql'>003_create_conta.sql</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/migrations/006_create_tipos_transferencia.sql'>006_create_tipos_transferencia.sql</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/migrations/005_create_pessoa_juridica.sql'>005_create_pessoa_juridica.sql</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
								<tr style='border-bottom: 1px solid #eee;'>
									<td style='padding: 8px;'><b><a href='https://github.com/ecbDeveloper/go-money/blob/master/internal/db/migrations/tern.conf'>tern.conf</a></b></td>
									<td style='padding: 8px;'>Code>❯ REPLACE-ME</code></td>
								</tr>
							</table>
						</blockquote>
					</details>
				</blockquote>
			</details>
		</blockquote>
	</details>
</details>

---

## Getting Started

### Prerequisites

This project requires the following dependencies:

- **Programming Language:** Go
- **Package Manager:** Go modules

### Installation

Build go-money from the source and intsall dependencies:

1. **Clone the repository:**

    ```sh
    ❯ git clone https://github.com/ecbDeveloper/go-money
    ```

2. **Navigate to the project directory:**

    ```sh
    ❯ cd go-money
    ```

3. **Install the dependencies:**

<!-- SHIELDS BADGE CURRENTLY DISABLED -->
	<!-- [![go modules][go modules-shield]][go modules-link] -->
	<!-- REFERENCE LINKS -->
	<!-- [go modules-shield]: https://img.shields.io/badge/Go-00ADD8.svg?style={badge_style}&logo=go&logoColor=white -->
	<!-- [go modules-link]: https://golang.org/ -->

	**Using [go modules](https://golang.org/):**

	```sh
	❯ go build
	```
--- 

### Usage

Run the project with:

**Using [go modules](https://golang.org/):**
```sh
go run {entrypoint}
```

<div align="right">

[![][back-to-top]](#top)

</div>


[back-to-top]: https://img.shields.io/badge/-BACK_TO_TOP-151515?style=flat-square


---
