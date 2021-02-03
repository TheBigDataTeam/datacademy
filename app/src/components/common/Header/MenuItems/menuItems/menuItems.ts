interface Item {
    type: 'item';
    title: string;
    url: string;
  }
  
  export const menuItems: Item[] = [
    {
      type: 'item',
      title: 'Project',
      url: '/project',
    },
    {
      type: 'item',
      title: 'Courses',
      url: '/courses',
    },
    {
      type: 'item',
      title: 'Authors',
      url: '/authors',
    },
    {
      type: 'item',
      title: 'FAQ',
      url: '/faq',
    },
  ];
  