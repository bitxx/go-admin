import { Button, ButtonProps } from "antd";
import { ButtonType } from "antd/es/button";
import React, { useState } from "react";

// 定义 LoadingButton 的参数类型
interface LoadingButtonProps extends Omit<ButtonProps, "onClick"> {
	onClick?: (done: () => void, event: React.MouseEvent<HTMLElement>) => void;
	type?: ButtonType;
}

const LoadingButton: React.FC<LoadingButtonProps> = ({ onClick, children, type, ...restProps }) => {
	const [loading, setLoading] = useState(false);

	const handleClick = (event: React.MouseEvent<HTMLElement>) => {
		setLoading(true); // 开始 loading
		const done = () => setLoading(false); // 提供回调函数，供外部控制结束 loading

		try {
			if (onClick) {
				onClick(done, event); // 执行传入的逻辑，提供 `done` 作为参数
			}
		} catch (error) {
			done(); // 捕获异常时也结束 loading
		}
	};

	// 保留传入的 key 值，并通过 ...restProps 传递其他 Button 的属性
	return (
		<Button {...restProps} loading={loading} type={type} onClick={handleClick}>
			{children}
		</Button>
	);
};

export default LoadingButton;
