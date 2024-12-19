import { Switch } from "antd";
import React, { useState } from "react";

interface LoadingSwitchProps {
  checked: boolean;
  checkedChildren: React.ReactNode;
  unCheckedChildren?: React.ReactNode;
  onChange: (checked: boolean) => Promise<void>;
}

const LoadingSwitch: React.FC<LoadingSwitchProps> = ({ checked, onChange, checkedChildren, unCheckedChildren, ...restProps }) => {
  const [loading, setLoading] = useState(false);

  const handleChange = async (checked: boolean) => {
    setLoading(true);
    try {
      await onChange(checked); // 调用父组件传入的异步变更函数
    } finally {
      setLoading(false);
    }
  };

  return (
    <Switch
      {...restProps}
      loading={loading}
      checked={checked}
      onChange={handleChange}
      checkedChildren={checkedChildren}
      unCheckedChildren={unCheckedChildren}
    />
  );
};

export default LoadingSwitch;
