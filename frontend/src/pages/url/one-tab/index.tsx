import { Card, List, Image } from 'antd';
import { PageContainer } from '@ant-design/pro-layout';
import { useRequest } from 'umi';
import { queryFakeList } from './service';

const CardList = () => {
  const { data } = useRequest(() => {
    return queryFakeList();
  });

  const list = data || [];

  return (
    <PageContainer>
      <div>
        <List
          grid={{
            gutter: 16,
            xs: 1,
            sm: 2,
            md: 4,
            lg: 4,
            xl: 6,
            xxl: 6,
          }}
          dataSource={list}
          renderItem={item => (
            <List.Item>
              <Card title={item.title}>
                <List
                  dataSource={item.urls}
                  renderItem={itemUrl => (
                    <List.Item>
                      <Image
                        width={16}
                        height={16}
                        src={itemUrl.icon}
                      />
                      <a href={itemUrl.url} target="_blank">{itemUrl.title}</a>
                    </List.Item>
                  )}
                >
                </List>
              </Card>
            </List.Item>
          )}
        />,
      </div>
    </PageContainer>
  );
};

export default CardList;
