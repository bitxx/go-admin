import React, { Suspense } from "react";
import { PageLoader } from "../Loading";

// 将 LazyComponent 类型更加明确化
const LazyComponent = (Comp: React.LazyExoticComponent<React.ComponentType<any>>) => {
  return (
    <Suspense fallback={<PageLoader />}>
      <Comp />
    </Suspense>
  );
};

export default LazyComponent;
