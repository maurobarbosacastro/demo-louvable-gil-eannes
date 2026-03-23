import {Injectable} from "@angular/core";
import {rewardDetails, rewards, rewardSummary} from "@app/mock-api/common/rewards/data";
import {FuseMockApiService} from "@fuse/lib/mock-api";
import {cloneDeep} from "lodash-es";
import {Reward, RewardDetails, RewardSummary} from "@app/modules/client/rewards/models/reward.interface";

@Injectable({ providedIn: 'root' })
export class RewardsMockApi {
    private _rewards: Reward[] = rewards;
    private _rewardSummary: RewardSummary = rewardSummary;
    private _rewardDetails: RewardDetails[] = rewardDetails;

    /**
     * Constructor
     */
    constructor(private _fuseMockApiService: FuseMockApiService) {
        // Register Mock API handlers
        this.registerHandlers();
    }

    // -----------------------------------------------------------------------------------------------------
    // @ Public methods
    // -----------------------------------------------------------------------------------------------------

    /**
     * Register Mock API handlers
     */
    registerHandlers(): void {
        // -----------------------------------------------------------------------------------------------------
        // @ rewards - GET
        // -----------------------------------------------------------------------------------------------------
        this._fuseMockApiService
            .onGet('api/ms-tagpeak/rewards')
            .reply(({ request }) => {
                const status: string[] = [request.params.get('status')];
                const statusFormatted: string[] = status?.[0]?.split(',');

                // Clone the rewards
                let rewards: Reward[] = cloneDeep(this._rewards);

                if (statusFormatted?.length > 0) {
                    rewards = rewards.filter((reward: Reward) => statusFormatted.includes(reward.status))
                }

                const totalSize: number = rewards.length;

                return [200, { rewards: rewards, totalSize: totalSize }];
            });
        // -----------------------------------------------------------------------------------------------------
        // @ rewardSummary - GET
        // -----------------------------------------------------------------------------------------------------
        this._fuseMockApiService
            .onGet('api/ms-tagpeak/reward/summary')
            .reply(() =>{
                    // Clone the rewardSummary
                    const rewardSummary: RewardSummary = cloneDeep(this._rewardSummary);

                    return [200, rewardSummary];
                }
            );

        // -----------------------------------------------------------------------------------------------------
        // @ rewardById - GET
        // -----------------------------------------------------------------------------------------------------
        this._fuseMockApiService
            .onGet(`api/ms-tagpeak/reward`)
            .reply(({ request }) => {
                    const id: string = request.params.get('id');
                    // Clone the reward
                    const reward: RewardDetails = cloneDeep(this._rewardDetails.find(reward => reward.uuid === id));

                    return [200, reward];
                }
            );

        // -----------------------------------------------------------------------------------------------------
        // @ updateReward - PATCH
        // -----------------------------------------------------------------------------------------------------
        this._fuseMockApiService
            .onPatch(`api/ms-tagpeak/reward`)
            .reply(({ request }) => {
                    const id: string = request.body.id;
                    const status = cloneDeep(request.body.status);

                    const reward: Reward = this._rewards.find(reward => reward.uuid === id);
                    reward.status = status;
                    reward.stopDate = null;

                    const rewardDetails: RewardDetails = this._rewardDetails.find(reward => reward.uuid === id);
                    rewardDetails.status = status;
                    rewardDetails.stopDate = null;

                    return [200, reward];
                }
            );
    }
}
