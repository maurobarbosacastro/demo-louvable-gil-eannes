export interface DashboardInterface {
	cashbackSection: CashbackSection;
	indicatorsSection: IndicatorsSection;
}

export interface CashbackSection {
	totalValidatedCashbacks: number;
	totalStoppedCashbacks: number;
	totalPaidCashbacks: number;
	totalRequestedCashbacks: number;
}

export interface IndicatorsSection {
	totalUsers: TotalUsers;
	activeUsers: ActiveUsers;
	numTransactions: NumTransactions;
	totalGMV: TotalGMV;
	averageTransactionAmount: AverageTransactionAmount;
	totalRevenue: TotalRevenue;
}

export interface TotalUsers {
	allTime: number;
	lastMonth: number;
	compareLastMonth: string;
	percentageChange: number;
}

export interface ActiveUsers {
	allTime: number;
	last12Months: number;
	compareLastMonth: string;
	percentageChange: number;
}

export interface NumTransactions {
	allTime: number;
	currentMonth: number;
	compareLastMonth: string;
	percentageChange: number;
}

export interface TotalGMV {
	allTime: number;
	currentMonth: number;
	compareLastMonth: string;
	percentageChange: number;
}

export interface AverageTransactionAmount {
	allTime: number;
	currentMonth: number;
	compareLastMonth: string;
	percentageChange: number;
}

export interface TotalRevenue {
	allTime: number;
	currentMonth: number;
	compareLastMonth: string;
	percentageChange: number;
}

export interface StatisticsByMonthInterface {
	[key: string]: MonthCountsSection;
}

export interface MonthCountsSection {
	totalUsers: number;
	activeUsers: number;
	numTransaction: number;
	totalGMV: number;
	avgTransactionAmount: number;
	totalRevenue: number;
}

export interface TransactionsStatusInterface {
	[key: string]: TransactionStatus;
}

export interface TransactionStatus {
	status: string;
	count: number;
	warning: number;
	value: number;
}

export interface RewardByCurrenciesInterface {
	[key: string]: RewardCurrenciesInterface;
}

export interface RewardCurrenciesInterface {
	[key: string]: RewardByCurrencies;
}

export interface RewardByCurrencies {
	currency: string;
	state: string;
	totalRewards: number;
}
