import { GithubOutlined } from '@ant-design/icons';
import { DefaultFooter } from '@ant-design/pro-components';
import React from 'react';

const Footer: React.FC = () => {
  return (
    <DefaultFooter
      style={{
        background: 'none',
      }}
      copyright="Powered by Apply Admin"
      links={[
        {
          key: 'github',
          title: <GithubOutlined />,
          href: 'https://github.com/lycarrot/apply-admin',
          blankTarget: true,
        },
        {
          key: 'Apply Admin',
          title: 'Apply Admin',
          href: 'https://github.com/lycarrot/apply-admin',
          blankTarget: true,
        },
      ]}
    />
  );
};

export default Footer;
