import "./Aside.css"
import type React from 'react';

const asideVariants = ['comment', 'definition', 'result', 'example'] as const;

interface Props {
	type: (typeof asideVariants)[number];
	title: string;
	name?: string,
	children: React.ReactNode;
}

const Aside = ({ name, type, title, children }: Props) => {
	return (
		<aside aria-label={title} className={`starlight-aside starlight-aside--${type}`}>
			<p className="starlight-aside__title" aria-hidden="true">
				{title}
				{(name != '') && <span>{name}</span>}
			</p>
			<section className="starlight-aside__content">
				{children}
			</section>
		</aside>


	)
}

export default Aside
