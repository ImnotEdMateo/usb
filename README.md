# Upload Bullshit in my Server

This is a simple and lightweight self-hosted file uploader built with GoLang. UBS allows you to upload files temporarily to your server with a straightforward interface and no JavaScript requirements. Perfect for quick file sharing or temporary storage.

## Features

- **Temporary file storage**: Uploaded files are deleted automatically after 1 hour.
- **JavaScript-free**: Fully functional without requiring JavaScript on the client side.
- **Unique file routes**: Files are accessible via unique, random routes.
- **File size limits**: Enforces a maximum file size for uploads.
- **Secure uploads**: Blocks potentially dangerous file types (e.g., `.exe`, `.sh`, `.bat`).
- **Easy to self-host**: 0 external dependencies, runs on HTTP.

## Deployment

### Using Docker (Recommended)

1. Clone the repository:
   ```bash
   git clone https://github.com/ImnotEdMateo/ubs.git
   cd ubs
   ```

2. Create or modify the `.env` file with the following content:
   ```bash
   UBS_PORT=[your_port]
   ```

3. Build and start the container using `docker-compose`:
   ```bash
   docker-compose up --build -d
   ```

4. The application will be available at `[your_vps_ip]:<port>`. 

> [!NOTE]  
> To access the application directly from `[your_vps_ip]:<port>` in the Dockerfile, you will need to hardcode the port. If you'd prefer to avoid hardcoding, you can use a reverse proxy like Apache or Nginx to manage access to the port.

### Without Docker

1. Clone the repository:
   ```bash
   git clone https://github.com/ImnotEdMateo/ubs.git
   cd ubs
   ```

2. Build the application:
   ```bash
   export UBS_PORT=[your_port]
   go build -o ubs
   ```

3. Run the server:
   ```bash
   ./ubs
   ```

The server will run on `http://[your_vps_ip]:<your_port>` by default.

## Configuration

- **Upload directory**: Files are stored in the `uploads/` directory. Make sure this directory exists and is writable by the server.
- **File size limit and duration**: These can be configured in `config/settings.go`. Adjust the values as needed for your use case.

## Usage

1. Access the uploader interface at `http://localhost:[your_port]`.
2. Select a file to upload and click "UPLOAD."
3. After a successful upload, you'll be redirected to the unique route where the file can be downloaded.

## Security

- Files are validated to ensure they are not executable or potentially harmful.
- MIME type and extension checks are implemented to improve safety.
- Temporary files are automatically deleted after 1 hour to minimize storage use.

## Limitations

- UBS is designed for small-scale use and does not include advanced features like user authentication.
- Files are accessible via public routes; handle sensitive data with caution.

## TO DO

- [ ] Implement configurable themes from `config/settings.go`
- [ ] Optimize file validation logic.
- [ ] Less shitty code

---

Enjoy uploading your bullshit!
