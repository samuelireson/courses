import * as React from 'react';
import Giscus from '@giscus/react';

const id = 'inject-comments';

const Comments = () => {
	const [mounted, setMounted] = React.useState(false);

	React.useEffect(() => {
		setMounted(true);
	}, []);

	return (
		<div id={id}>
			{mounted ? (
				<Giscus
					category="Announcements"
					categoryId="DIC_kwDOM7U0wM4CjJEn"
					emitMetadata="0"
					id={id}
					inputPosition="top"
					lang="en"
					loading="lazy"
					mapping="pathname"
					reactionsEnabled="1"
					repoId="R_kgDOM7U0wA"
					repo="samuelireson/courses"
					theme="dark"
				/>
			) : null}
		</div>
	);
};

export default Comments;

