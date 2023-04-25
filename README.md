# EMIVNTestTask
Реализовать телеграмм-бот имеющий:
 ⁃ 6 разных сущностей<br/>
 ⁃ Администратор<br/>
 ⁃ Сёгун<br/>
 ⁃ Даймё<br/>
 ⁃ Самурай<br/>
 ⁃ Инкассатор<br/>
 ⁃ Карты<br/>


 ⁃ Администратор<br/>
 ⁃ Умеет создавать каждый вид сущности<br/>
 ⁃ Умеет создавать карты<br/>
 ⁃ Умеет привязывать подчиненных
( Сёгун - Дайме / Даймё - Самурай и тд )<br/>
 ⁃ Умеет привязывать карты к Даймё<br/>
 ⁃ Может посмотреть информацию по каждой сущности ( Общая информация + последние данные за смену )<br/>


 ⁃ Сёгун<br/>
 ⁃ Видит список своих подчиненных<br/>
(Даймё) и может через них провалиться в список их подчиненных (Самураи)<br/>
 ⁃ Может создавать карты и привязывать их к Даймё<br/>
 ⁃ Может просмотреть информацию по каждому подчиненному (Дайме - Общая информация + оборот самураев за смену / Самурай - Оборот конкретно этого самурая за смену )<br/>


 ⁃ Дайме<br/>
 ⁃ Видит список своих карт<br/>
 ⁃ Создает заявки на пополнение своих карт<br/>
 ⁃ Видит список своих самураев и может провалиться в каждого из них и посмотреть оборот за смену<br/>
 ⁃ Вводит остаток на картах на конец смены ( с 8:00 до 12:00 )<br/>
 ⁃ Привязывается к одному Сёгуну<br/>


 ⁃ Самурай<br/>
 ⁃ Вводит оборот за смену (с 8:00 до
12:00)<br/>
 ⁃ Привязывается к одному Даймё<br/>


 ⁃ Инкассатор<br/>
 ⁃ Выполняет заявки Дайме на пополнение карт<br/>


 ⁃ Карты<br/>
 ⁃ Привязываются к даймё<br/>
 ⁃ Содержат информацию о банке эмитенте<br/>
 ⁃ Суточных лимитах ( По умолчанию 2 000 000 )<br/>
 
 - Инструкция к использованию<br/>
 <br/>
 - Администратор<br/>
 - admin create_card [cardID] [bankInfo] {LimitInfo}<br/>
 - admin connect_card [cardID] [owner]<br/>
 - admin create_shogun [Nickname] [TG username]<br/>
 - admin create_daimyo [Nickname] [TG username]<br/>
 - admin create_samurai [Nickname] [TG username]<br/>
 - admin create_collector [Nickname] [TG username]<br/>
 - admin set_daimyo_owner [daimyo nickname] [shogun nickname]<br/>
 - admin set_samurai_owner [samurai nickname] [daimyo nickname]<br/>
 - admin get_shogun_info [shogunID]<br/>
 - admin get_daimyo_info [daimyoID]<br/>
 - admin get_samurai_info [samuraiID]<br/>
 - admin get_collector_info [collectorID]<br/>
 <br/>
 - Сёгун<br/>
 - shogun daimyos - Просмотр подчиненных(Даймё)<br/>
 - shogun samurais [daimyo nickname] - Просмотр подчиненных самураев у конкретного даймё<br/>
 - shogun create [cardID] [bankInfo] {LimitInfo}<br/>
 - shogun connect [cardID] [owner] - Привезка карты к даймё по нику<br/>
 <br/>
 - Дайме<br/>
 - daimyo samurais<br/>
 - daimyo set [cardID] [balance] - Остаток на карте<br/>
 - daimyo cards<br/>
 - daimyo application [cardID] [value] - Создание заявки на пополнение карты[cardID] до суммы [value]<br/>
 <br/>
 - Самурай<br/>
 - samurai turnover [value] - оборот за смену<br/>
 <br/>
 - Инкассатор<br/>
 - collector show - Показать запросы на пополнение<br/>
 - collector apply [cardID] [value] - Выполнить апрос на пополнение<br/>
