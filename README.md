## Descripción

QvaRate API es un servicio que proporciona la tasa de cambio del Peso Cubano (CUP) frente a las principales monedas manejadas en Cuba (USD, EUR y MLC) dentro de un rango de fechas especificado. Adicionalmente, permite la exportación de estos datos en formato Excel.

---

## Endpoints

### 1. Obtener tasas de cambio en un rango de fechas

**URL:**

```
GET https://qvarateapi.onrender.com/api/get-currency/{startdate}/{enddate}
```

**Descripción:**  
Este endpoint devuelve la tasa de cambio del CUP para USD, EUR y MLC dentro del rango de fechas especificado.

**Parámetros:**

- `{startdate}`: Fecha de inicio en formato `YYYY-MM-DD`.
- `{enddate}`: Fecha de fin en formato `YYYY-MM-DD`.

**Ejemplo de solicitud:**

```
GET https://qvarateapi.onrender.com/api/get-currency/2025-02-17/2025-02-19
```

**Ejemplo de respuesta:**

```json
{
    "result": [
        {
            "date": "2025-02-19T00:00:00Z",
            "usd": 340,
            "eur": 345,
            "mlc": 265
        },
        {
            "date": "2025-02-18T00:00:00Z",
            "usd": 340,
            "eur": 345,
            "mlc": 260
        },
        {
            "date": "2025-02-17T00:00:00Z",
            "usd": 340,
            "eur": 345,
            "mlc": 255
        }
    ]
}
```

---

### 2. Exportar datos en formato Excel

**URL:**

```
GET https://qvarateapi.onrender.com/api/get-excel/{startdate}/{enddate}
```

**Descripción:**  
Este endpoint genera y permite la descarga de un archivo Excel con los valores de la tasa de cambio del CUP frente a USD, EUR y MLC en el rango de fechas especificado.

**Parámetros:**

- `{startdate}`: Fecha de inicio en formato `YYYY-MM-DD`.
- `{enddate}`: Fecha de fin en formato `YYYY-MM-DD`.

**Ejemplo de solicitud en JavaScript:**

```javascript
const handleExportExcel = () => {
  const startdate = '2024-01-07';
  const enddate = '2025-01-17';

  const url = `https://qvarateapi.onrender.com/api/get-excel/${startdate}/${enddate}`;

  fetch(url, {
    method: 'GET',
    headers: {
      'Accept': 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet',
    },
    mode: 'cors',
  })
    .then(response => {
      if (!response.ok) {
        throw new Error(`Error al descargar el archivo: ${response.status} ${response.statusText}`);
      }
      return response.blob();
    })
    .then(blob => {
      const a = document.createElement('a');
      const fileUrl = URL.createObjectURL(blob);
      a.href = fileUrl;
      a.download = `Tasa de Cambio ${startdate} - ${enddate}.xlsx`;
      a.click();

      URL.revokeObjectURL(fileUrl);
      console.log('Archivo descargado exitosamente');
    })
    .catch(error => {
      console.error('Hubo un error:', error.message);
      alert(`Hubo un error al descargar el archivo: ${error.message}`);
    });
};
```

---

## Consideraciones

- **Formato de fecha:** Todas las fechas deben estar en formato `YYYY-MM-DD`.
- **CORS:** La API admite solicitudes CORS para facilitar su integración con aplicaciones web.
- **Monedas disponibles:**
    - **USD:** Dólar Estadounidense
    - **EUR:** Euro
    - **MLC:** Moneda Libremente Convertible (MLC)
