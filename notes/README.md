# notes

## Adding content

Each new chapter should be added to the `chapter` directory, following the naming convention which is found there. These files can then be included into the `master.tex` file found in the relevant course folder. Do not number chapters explicitly in their filenames, since this places barriers to course rearrangement.

## Content tagging

The `preamble.tex` file defines macros for content tagging, which has especially nice output in the web version. Try to tag the content on the most granular level possible. For example, if a whole section constitutes basic knowledge, then it should be tagged `\basic`. It should never be the case that a `\challenging` concept is contained in a `\basic` section; instead label the subsections.

## Exercises

In order for the file format conversion to work effectively, it is important to be careful about how exercises are declared in the `.tex`. For example,

```tex
\section{Exercises}
\subsection{Exercises \basic}
\begin{exercise}
	\begin{problem}
        This is the problem statement.
	\end{problem}
	\begin{solution}
		This is the solution.
	\end{solution}
\end{exercise}
```

with the same layout for the corresponding `\intermediate` and `\challenging` exercises.

