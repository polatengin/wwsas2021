import React, { FunctionComponent, useEffect } from 'react';

import { ShoppingCartIcon, StoreLogoIcon, UserIcon } from './Icons';

export const App: FunctionComponent = () => {
  return (
    <div className="flex flex-col bg-gray-300 h-3/4 px-4 rounded-xl">
      <header className="flex my-4 items-center">
        <StoreLogoIcon className="w-6 h-6 mr-2" />
        <span className="text-lg">WorldWide Software Architecture Summit</span>
        <span className="flex-grow"></span>
        <ShoppingCartIcon className="w-6 h-6" />
        <UserIcon className="w-6 h-6" />
      </header>
      <main className="flex-grow">
        adasd
      </main>
      <footer className="my-4">
      </footer>
    </div>
  );
};
