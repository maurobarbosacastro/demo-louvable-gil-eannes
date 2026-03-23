import type {PaginationInterface, PaginationOptions} from '../interfaces/common.interface';
import type {CashbackTableResponse, ServiceResponse} from '../interfaces/dashbaord.interface';
import {BlockStack, Button, Card, IndexTable, Text, useIndexResourceState} from '@shopify/polaris';
import {useEffect, useState} from 'react';
import {useI18n} from '@shopify/react-i18n';
import {format, parseISO} from 'date-fns';
import {bulkEditTransaction, getCashbacks} from '../service/dashboard.service';
import {logout} from '../utils/auth.utils';

type Props = {
	cashback: PaginationInterface<CashbackTableResponse>;
	beUrl: string;
	boUrl: string;
	storeUuid: string;
	shop: string;
	setRefreshParent: React.Dispatch<React.SetStateAction<number>>;
};

export default function Cashback({cashback, storeUuid, shop, setRefreshParent}: Props) {
  const [i18n] = useI18n();
  const [data, setData] = useState<CashbackTableResponse[]>(cashback.data.map( d => ({...d, id: d.uuid})));
  const [pagination, setPagination] = useState<PaginationOptions>({
    page: 0,
    sort: 'date desc',
    limit: 30
  });
  const [refreshTrigger, setRefreshTrigger] = useState(0);
  const resourceName = {
    singular: 'order',
    plural: 'orders',
  };
  const {selectedResources, allResourcesSelected, handleSelectionChange, clearSelection} =
    useIndexResourceState(data);
	const [isLoading, setIsLoading] = useState(false);

  const update = (state: string) => {
    bulkEditTransaction({
      uuids: selectedResources,
      state: state,
    })
      .then( (res) => {
        if (res.status === 401){
          logout(shop)
	        return;
        }

				clearSelection();
        setData(data.map( d => {
          if (d.uuid === selectedResources[0]) {
            return {...d, status: state};
          }
          return d;
        }));
      })
  }

  const promotedBulkActions = [
    {
      content: 'Confirm reward',
      onAction: () => {
        update('VALIDATED')
      },
    },
    {
      content: 'Cancel reward',
      destructive: true,
      onAction: () => {
        update('REJECTED')
      },
    },
  ];

  useEffect(() => {
    getCashbacks(storeUuid, pagination.limit, pagination.page, pagination.sort)
      .then((res: ServiceResponse<PaginationInterface<CashbackTableResponse>>) => {
        if (res.status === 401){
	        setIsLoading(false);
          logout(shop)
        }

        if (res.data){
          setData(res.data.data.map( d => ({...d, id: d.uuid})));
          setPagination({
            page: res.data.page,
            limit: res.data.limit,
            sort: res.data.sort
          });
	        setIsLoading(false);
        }
      })
  }, [pagination.page, pagination.limit, pagination.sort, refreshTrigger]);


  return (
    <Card>
      <BlockStack gap="1000">
        <div style={{display: 'flex', justifyContent: 'space-between'}}>
          <Text as="h3" variant="headingLg" fontWeight="medium" >
            Orders
          </Text>

          <Button
            size="large"
            onClick={ () => {
	            setRefreshParent(refreshTrigger)
							setIsLoading(true);
              setRefreshTrigger(prev => prev + 1);
              setPagination({
                page: 0,
                limit: 30,
                sort: 'date desc'
              });
            }}
            loading={isLoading}
            disabled={isLoading}
          >
            Refresh
          </Button>
        </div>

        <IndexTable
          resourceName={resourceName}
          itemCount={data.length}
          selectedItemsCount={
            allResourcesSelected ? 'All' : selectedResources.length
          }
          promotedBulkActions={promotedBulkActions}

          onSelectionChange={handleSelectionChange}
          headings={[
            {title: 'Order'},
            {title: 'Date'},
            {title: 'Customer'},
            {title: 'Total', alignment: 'end'},
            {title: 'Tagpeak UUID'},
            {title: 'Status'}
          ]}
          pagination={{
            hasNext:  (pagination.page + 1) < cashback.totalPages,
            hasPrevious: pagination.page > 0,
            onNext: () => {
              setPagination({
                page: pagination.page + 1,
                limit: pagination.limit,
                sort: pagination.sort
              });
            },
            onPrevious: () => {
              setPagination({
                page: pagination.page - 1,
                limit: pagination.limit,
                sort: pagination.sort
              });
            },
          }}
        >
          {
            data.map(
              (
                {id, uuid, amountSource, currencySource, date, email, exitId, status},
                index,
              ) => (
                <IndexTable.Row
                  id={id as string}
                  key={id as string}
                  selected={selectedResources.includes(uuid)}
                  position={index}
                  disabled={status === 'VALIDATED'}
                >
                  <IndexTable.Cell>
                    <Text variant="bodyMd" fontWeight="bold" as="span">
                      {exitId}
                    </Text>
                  </IndexTable.Cell>
                  <IndexTable.Cell>
                    {format(parseISO(date), "d MMM y 'at' HH:mm")}
                  </IndexTable.Cell>
                  <IndexTable.Cell>{email}</IndexTable.Cell>
                  <IndexTable.Cell>
                    <Text as="span" alignment="end" numeric>
                      {i18n.formatCurrency(amountSource, {
                        currency: currencySource,
                        form: 'short'
                      })}
                    </Text>
                  </IndexTable.Cell>
                  <IndexTable.Cell>{uuid}</IndexTable.Cell>
                  <IndexTable.Cell
                  >
                    {
                      status === 'TRACKED' &&
                      <Text as="span" tone={"subdued"}>
                        Pending
                      </Text>
                    }
                    {
                      status === 'VALIDATED' &&
                      <Text as="span" tone={"success"}>
                        Confirmed
                      </Text>
                    }
                    {
                      status === 'REJECTED' &&
                      <Text as="span" tone={"critical"}>
                        Rejected
                      </Text>
                    }
                  </IndexTable.Cell>
                </IndexTable.Row>
              )
            )
          }
        </IndexTable>
      </BlockStack>
    </Card>
  );
}
