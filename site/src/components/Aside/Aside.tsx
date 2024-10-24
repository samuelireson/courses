import { Icon } from '@astrojs/starlight/components';
import "./Aside.css"

const asideVariants = ['comment', 'definition', 'result', 'example'] as const;
const icons = { comment: 'comment', definition: 'pencil', result: 'information', example: 'magnifier' } as const;

interface Props {
	type: (typeof asideVariants)[number];
	title: string;
}

const Aside = ({ type, title }: Props) => {
	return (
		<aside aria-label={title} className={`starlight-aside starlight-aside--${type}`}>
			<p className="starlight-aside__title" aria-hidden="true">
				<Icon name={icons[type]} class="starlight-aside__icon" />{title}
			</p>
			<section className="starlight-aside__content">
				<slot />
			</section>
		</aside>


	)
}

export default Aside
