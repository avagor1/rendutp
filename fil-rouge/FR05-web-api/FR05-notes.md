# FR05 – Notes de la séance 1

## Étape 2 : requête et réponse

| Élément | Votre réponse |
|---|---|
| Méthode de la requête | `GET` |
| Chemin (sans les paramètres) | `/v1/forecast` |
| Paramètres de requête (noms) | `latitude`, `longitude`, `current` |
| Code de statut et son texte | `200 OK` |
| Valeur de l'en-tête `Content-Type` | `application/json` |

## Étape 3 : le JSON

Corps de la réponse reçu avec `curl.exe -i` :

```json
{"current":{"temperature_2m":8.4,"time":"2026-11-16T14:30","wind_speed_10m":12.6},"current_units":{"temperature_2m":"°C","time":"iso8601","wind_speed_10m":"km/h"},"latitude":48.89,"longitude":2.2,"timezone":"Europe/Paris"}
```

| Chemin dans le JSON | Exemple de valeur | Type JSON | Type Go |
|---|---|---|---|
| `latitude` | `48.89` | nombre | `float64` |
| `current` | `{"temperature_2m": 8.4, "time": "2026-11-16T14:30", "wind_speed_10m": 12.6}` | objet | `struct` |
| `current.time` | `"2026-11-16T14:30"` | chaîne | `string` |
| `current.temperature_2m` | `8.4` | nombre | `float64` |
| `current_units.temperature_2m` | `"°C"` | chaîne | `string` |

## Étape 4 : provoquer des erreurs

| Requête | Code | Corps | Famille | En faute |
|---|---|---|---|---|
| `/v1/forecast` (sans paramètre) | `400 Bad Request` | `{"error":true,"reason":"latitude et longitude obligatoires"}` | 4xx | ma requête (paramètres manquants) |
| `/v1/forecast?latitude=abc&longitude=2.20` | `400 Bad Request` | `{"error":true,"reason":"latitude et longitude doivent être des nombres valides"}` | 4xx | ma requête (latitude invalide) |
| `/v1/previsions` | `404 Not Found` | `404 page not found` (texte brut) | 4xx | ma requête (la route n'existe pas) |