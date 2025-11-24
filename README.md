# Spotify Web – Projet Go / GOHTML

Ce projet a pour objectif d’afficher des données récupérées depuis l’API **Spotify Web API** sur des pages web générées par un serveur HTTP écrit en Go, avec des templates **GOHTML**.

- Serveur HTTP en Go (port **8080**)
- Requêtes vers l’API Spotify
- Décodage du JSON (package `encoding/json`)
- Affichage des données via des templates HTML (GOHTML)
- Design libre du site (mise en page, couleurs, etc.)

---

## 🗂 Arborescence du projet

```text
.
├── requêtes/                 # (éventuels tests / exemples de requêtes)
├── siteWeb/
│   ├── assets/               # ressources statiques (CSS, images, JS…)
│   ├── controller/
│   │   ├── api.go            # logique d’appel à l’API Spotify
│   │   ├── apiController.go  # contrôleur dédié aux données Spotify
│   │   └── controller.go     # contrôleurs / handlers génériques
│   ├── router/
│   │   └── router.go         # définition des routes HTTP
│   └── template/
│       ├── damso.html        # page affichant les données pour Damso
│       ├── home.html         # page d’accueil
│       └── laylow.html       # page affichant les données pour Laylow
├── go.mod                    # module Go
└── main.go                   # point d’entrée du serveur
