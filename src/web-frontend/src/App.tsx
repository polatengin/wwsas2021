import React, { FunctionComponent, useEffect, useState } from 'react';

import { ShoppingCartIcon, StarIcon, StoreLogoIcon, TagIcon, UserIcon } from './Icons';

interface Product {
	Name: string
	Price: number
	InStock: boolean
	Rating: number
	ImageUrl: string
}

export const App: FunctionComponent = () => {
  const [productList, setProductList] = useState<Product[]>([]);

  useEffect(() => {
    const getProductList = async () => {
      const response = await fetch("/api/product/get-list");
      const data = await response.json();
      setProductList(data);
    };
    getProductList();
  }, []);

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
        <div className="bg-gray-300 rounded-xl p-4 mb-4 flex">
          <div className="flex-grow">
            <h1 className="text-2xl">
              <span className="block text-gray-800">Online Shopping</span>
              <span className="block text-green-600">Buy to be happier 🥳</span>
            </h1>
            <p className="my-4 text-base text-gray-500 text-lg">
              Anim aute id magna aliqua ad ad non deserunt sunt. Qui irure qui lorem cupidatat commodo. Elit sunt amet fugiat veniam occaecat fugiat aliqua.
            </p>
            <span className="inline-block cursor-pointer rounded-md shadow px-4 py-2 border border-transparent text-lg font-medium rounded-md text-white bg-green-600 hover:bg-green-700">
              Get started
            </span>
          </div>
          <img className="rounded-xl shadow" src="https://picsum.photos/300/200?color&random=2954" />
        </div>
        <div className="flex items-center">
          <TagIcon className="w-6 h-6 mr-2" />
          <span className="text-xl">Discounted Products</span>
          <span className="flex-grow"></span>
          <span>View all ▸</span>
        </div>
        <div>
          {productList.map((product, index) => {
            return (
              <div key={index}>
                asdasd
              </div>
            );
          })}
        </div>
      </main>
      <footer className="my-4">
      </footer>
    </div>
  );
};
