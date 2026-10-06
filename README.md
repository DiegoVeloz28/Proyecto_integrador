# Gestión Centralizada de Imágenes Corporativas - Proyecto Integrador

## Descripción del sistema
El sistema es una aplicación empresarial cliente-servidor enfocada en centralizar, estandarizar y asegurar el ciclo de vida de los activos de imagen corporativos. Desarrollado con un backend en Go y un servicio de autenticación en FastAPI, su objetivo es solucionar los problemas de trazabilidad e inconsistencia de información entre la Sede Matriz Norte y la Sede Operativa Sur de una empresa logística[cite: 1]. El sistema permite a los usuarios autenticados subir fotografías (como evidencias de entregas, inspecciones o campañas) asociadas a metadatos descriptivos, garantizando un almacenamiento seguro y alimentando un componente estadístico en tiempo real.

## Integrantes
* Flavio Granizo
* Diego Veloz
* Alex Gaibor

## Alcance preliminar
* **Gestión de imágenes:** Carga masiva o individual de archivos binarios de imagen (JPEG, PNG, WebP) adjuntando metadatos obligatorios como título, categoría, descripción y etiquetas.
* **Búsqueda y filtrado:** Consulta avanzada de imágenes mediante filtros cruzados (por fecha, categoría, sede o autor) que devuelven resultados paginados de forma eficiente.
* **Roles y permisos:** Soporte para "Usuario Estándar" (encargado de cargar y consultar imágenes) y "Administrador / Auditor" (encargado de auditar, eliminar imágenes obsoletas y ver métricas globales).
* **Seguridad:** Autenticación obligatoria e integración con FastAPI para la validación de sesiones mediante Bearer Tokens en cada solicitud, previniendo accesos no autorizados.
* **Panel estadístico:** Generación de métricas para visualizar indicadores clave (KPIs) como el volumen de almacenamiento consumido en MB/GB y la frecuencia mensual de uso por sede.

## Estructura inicial del repositorio
De acuerdo a la arquitectura modular establecida, el repositorio se divide principalmente entre el servicio de autenticación y el backend principal:


/
├── auth-service/          # Servicio externo encargado de la autenticación de usuarios mediante FastAPI
├── backend-go/            # Backend de alto rendimiento desarrollado en Go
│   ├── handlers/          # Manejadores para las peticiones HTTP
│   ├── logic/             # Lógica de negocio, validación de formatos, tamaños y metadatos
│   └── repositories/      # Repositorios para la conexión y consultas a la base de datos relaciona
├── docs/                  # Documentación, diagramas de casos de uso y matriz de trazabilidad
└── README.md              # Archivo principal de descripción del proyecto