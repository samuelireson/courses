# site

## Commands

All commands are run from the root of the project, from a terminal:

| Command                   | Action                                           |
| :------------------------ | :----------------------------------------------- |
| `npm install`             | Installs dependencies                            |
| `npm run dev`             | Starts local dev server at `localhost:4321`      |
| `npm run build`           | Build your production site to `./dist/`          |
| `npm run preview`         | Preview your build locally, before deploying     |
| `npm run astro ...`       | Run CLI commands like `astro add`, `astro check` |
| `npm run astro -- --help` | Get help using the Astro CLI                     |

## Workflow

The idea is that there is very little to touch in this directory as far as adding pages to a course. However, upon converting a `.tex` file into `.mdx` format for the site, it is necessary to tell astro that this new file exists and should be shown on the webpage. This can be done by adding to the `astro.config.mjs` file located in the root of this directory.

```js
{ label: 'Linear Algebra', slug: 'mlnn/linear-algebra' },
{ label: 'Statistics', slug: 'mlnn/statistics' },
+ { label: 'New Page Title', slug: 'course/newpage' },
```

## Changing content

Changes to the content of a course should always occur via the `.tex` file so that there is consistency between the webpage and `.pdf` formats.
