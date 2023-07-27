# features_practice

Данная go - библиотека, которая позволяет эффективно вычислять признаки по данным из базы данных ArangoDB.

[Пример](main/main.go) использования

Список всех методов для адресов:
- TotalGetAddr (сумма полученых BTC на адрес)
- BalanceAddr (баланс адреса)
- FirstTimeAddr (время первого появления адреса)
- LastTimeAddr (время последнего появления адреса)
- CountOutTx (количество входящих транзакций на адрес)
- CountInTx (количество исходящих транзакций от адреса)
- CountInAddr (количество адресов на которые уходили средства)
- CountOutAddr (количество адресов с которых приходили средства)
- CountSharedAddr (количество общих адресов среди countInAddr и countOutAddr)
- TotalCountAddr (общее количество адресов среди countInAddr и countOutAddr)
- CountUniqueAddr (количество уникальных адресов среди countInAddr и countOutAddr)
- AverageCountOutAddr (среднее количество адресов во входных транзакциях)
- AverageCountInAddr (среднее количество адресов в выходящих транзакциях)
- NmotifAddr (набор путей длины n из адреса A в адрес B)

Список всех методов для кластеров:
- TotalGetClust (сумма полученых BTC на кластер)
- BalanceClust (сумма балансов всех адресов из кластера)
- FirstTimeClust (время первого появления какого-либо адреса из кластера)
- LastTimeClust (время последнего появления какого-либо адреса из кластера)
- CountOutTxClust (количество входящих транзакций на кластер)
- CountInTxClust (количество исходящих транзакций от кластера)
- CountInClust (количество адресов на которые уходили средства)
- CountOutClust (количество адресов с которых приходили средства)
- CountSharedClust (количество общих адресов среди countInClust и countOutClust)
- TotalCountClust (общее количество адресов среди countInClust и countOutClust)
- CountUniqueClust (количество уникальных адресов среди countInClust и countOutClust)
- AverageCountOutClust (среднее количество адресов во всех входных транзакциях)
- AverageCountInClust (среднее количество адресов во всех выходящих транзакциях)
- NmotifClust (набор путей длины n из кластера A в кластер B)
