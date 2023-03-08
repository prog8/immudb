window.onload = function() {
    //<editor-fold desc="Changeable Configuration Block">
  
    // the following lines will be replaced by docker/configurator, when it runs in a docker-container
    window.ui = SwaggerUIBundle({
      urls: [
        {"url": "/api/docs/authorizationschema.swagger.json", "name": "Authorization API"},
        {"url": "/api/docs/documentsschema.swagger.json", "name": "Documents API"},
        {"url": "/api/docs/schema.swagger.json", "name": "KV and SQL API"}
      ],
      dom_id: '#swagger-ui',
      deepLinking: true,
      presets: [
        SwaggerUIBundle.presets.apis,
        SwaggerUIStandalonePreset
      ],
      plugins: [
        SwaggerUIBundle.plugins.DownloadUrl
      ],
      layout: "StandaloneLayout"
    });
  
    //</editor-fold>
  };
  