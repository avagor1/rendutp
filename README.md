# CampusLink

Interface web statique de gestion des **salles**, des **équipements** et des **incidents** d'un campus.
Projet réalisé dans le cadre du module **FR01 HTML / CSS** (octobre 2026).

Uniquement du **HTML5 et du CSS3** : aucun JavaScript, aucune bibliothèque, aucun framework.

---

## Lancer le projet

1. Télécharger ou cloner le dossier `campuslink`.
2. Ouvrir `index.html` dans un navigateur (double-clic), **ou** dans Visual Studio Code avec l'extension **Live Server** (clic droit → *Open with Live Server*).

Aucune installation n'est nécessaire : les liens sont relatifs, le site fonctionne en local.

Navigateurs recommandés : versions récentes de Chrome, Edge, Safari ou Firefox (le site utilise `:has()`, `color-mix()` et `backdrop-filter`).

---

## Arborescence

```
campuslink/
├── index.html               # Tableau de bord (chiffres clés, incidents prioritaires)
├── salles.html              # Les 122 salles, regroupées par étage, avec leur disponibilité
├── equipements.html         # Tableaux des vidéoprojecteurs et des prises de chaque salle
├── incidents.html           # Tableau des incidents
├── incident-detail.html     # Fiche détaillée d'un incident (INC-042)
├── declarer-incident.html   # Formulaire de déclaration d'un incident
├── css/
│   └── style.css            # Feuille de style unique
└── README.md
```

---

## Les pages

| Page | Contenu |
|---|---|
| Accueil | 4 chiffres clés, incidents prioritaires, salles concernées, bouton « Déclarer un incident » |
| Salles | 122 cartes (Amphi Bleu, Souk, salles 101 à 620) regroupées par étage : nombre de places, badge de disponibilité, liseré coloré selon l'état |
| Équipements | 2 tableaux (vidéoprojecteurs, prises électriques) : une ligne par salle avec son état |
| Incidents | Tableau Réf. / Titre / Lieu / Gravité / Statut, lien vers le formulaire |
| Détail d'un incident | Description, historique daté en frise verticale, bloc d'informations (lieu, équipement, statut, déclarant) |
| Déclarer un incident | Formulaire : titre, salle, équipement, gravité, description, e-mail |

Le bouton **« Déclarer un incident »** est présent sur toutes les pages sauf le formulaire lui-même.

Les salles proposées dans le formulaire sont les mêmes que celles de la page Salles et des tableaux d'équipements : Amphi Bleu, 101 à 120, 201 à 220, Souk, puis 301 à 620.

Parcours principal : **Accueil → Incidents → Détail**, ou **Accueil → Déclarer un incident → Incidents**.

---

## Choix techniques

### HTML

- Même structure sur chaque page : `header`, `main#contenu` (titre de page, `nav.navigation`, contenu), `footer`.
- Navigation : le lien de la page en cours porte `aria-current="page"` (visuel + lecteurs d'écran).
- Balises choisies selon le sens : `table` avec `caption`, `thead` et `th scope` ; `dl` pour les paires étiquette / valeur ; `ol` et `time datetime` pour l'historique ; `article` et `aside` pour la fiche d'incident ; `fieldset` et `legend` pour la gravité ; une `section` par étage sur la page Salles.
- Titres : un seul `h1` par page (le nom du site), un `h2` pour le titre de la page, des `h3` pour les sections (par exemple chaque étage) et des `h4` pour les cartes de salles.
- Formulaire : chaque champ a un `label` relié par `for` / `id`, validation native (`required`, `type="email"`), aide associée par `aria-describedby`.
- Tableaux dans un conteneur `role="region"` focusable au clavier, qui défile horizontalement sur mobile.

### CSS

- Une seule feuille de style : variables → base → en-tête → navigation → composants → tableaux → historique → formulaire → pied de page → animations → responsive → mode sombre.
- **Variables CSS** (`--bleu`, `--ambre`, `--surface`, `--border`…) : les couleurs se changent en un seul endroit.
- Composants réutilisables : `card`, `badge` (variantes `badge-ok`, `badge-warn`, `badge-crit`, `badge-info`), `lien-bouton`, `liste`, `stat`, `meta`, `timeline`, `table-wrap`.
- Mise en page : **Flexbox** (navigation, listes, badges), **Grid** (cartes, chiffres clés, colonnes, formulaire).
- **Mobile first** : base en une colonne, puis media queries `min-width: 48rem` (2 colonnes) et `min-width: 64rem` (3 à 4 colonnes, mise en page à deux colonnes).
- Interface : en-tête en dégradé avec halos lumineux, navigation collante (`position: sticky`) avec effet de flou, cartes avec ombre et soulèvement au survol, boutons en forme de pilule.
- Couleur selon l'état : avec `:has()`, le liseré d'une carte et le trait à gauche d'une ligne de tableau reprennent la couleur du badge qu'elles contiennent (vert, ambre, rouge, bleu).
- Formulaire : choix de gravité présentés comme des cases cliquables qui prennent la couleur choisie, halo de focus, champ invalide signalé en rouge (`:user-invalid`).
- Animations : apparition progressive des blocs au chargement, pastille qui pulse sur les badges critiques.
- **Mode sombre** automatique (`prefers-color-scheme: dark`).
- Accessibilité : contrastes élevés, `:focus-visible` marqué, états signalés par un texte en plus de la couleur (les badges ne reposent pas sur la couleur seule), animations désactivées si l'utilisateur demande un mouvement réduit (`prefers-reduced-motion`).

---

## Charte graphique

| Couleur | Code | Utilisation |
|---|---|---|
| Bleu nuit | `#17365d` | En-tête, titres, boutons, liens |
| Bleu moyen | `#24508a` | Dégradés, focus, accents |
| Ambre | `#f2a900` | Liseré de l'en-tête et du pied de page, soulignement des titres, focus clavier |
| Gris très clair | `#f5f7fa` | Fond des pages |
| Vert | `#166534` sur `#d9f0e0` | Disponible, en service |
| Ambre foncé | `#7a4300` sur `#fdecc8` | À vérifier, partielle, moyen |
| Rouge | `#a61b1b` sur `#fbdcdc` | Hors service, indisponible, critique |
| Bleu info | `#1d4f7a` sur `#dbe9f6` | Occupée, faible |

Les liserés des cartes et des lignes de tableau utilisent des versions plus vives de ces couleurs (`--ok`, `--warn`, `--crit`, `--info`).

Police : **Arial, sans-serif**.

---

## Limites et suites prévues

Le site est **statique** : les données sont écrites en dur (états des salles et des équipements, nombre de places) et le formulaire n'enregistre rien (il renvoie vers la liste des incidents).

- [ ] JavaScript pour les filtres et la recherche (par étage, par état)
- [ ] Backend (Go) et base de données SQL pour enregistrer les incidents et mettre à jour les états
- [ ] Une fiche de détail par incident (toutes les lignes pointent aujourd'hui vers INC-042)
- [ ] Ajouter les autres équipements du formulaire (poste enseignant, micro sans fil)

---

## Auteurs

Projet réalisé par Valérie, Mbayang et Norah, Ynov, module **FR01 HTML / CSS**, octobre 2026.