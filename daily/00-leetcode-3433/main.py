from typing import List

class Solution:
    def countMentions(self, numberOfUsers: int, events: List[List[str]]) -> List[int]:
        offline_blocks = [[] for _ in range(numberOfUsers)]
        mentions = [0 for _ in range(numberOfUsers)]
        # process events
        for event in events:
            ename, etimestamp_str, eids_str = event[0], event[1], event[2]
            if ename == "OFFLINE":
                etimestamp = int(etimestamp_str)
                eid = int(eids_str)  # format: id123123
                user_block = offline_blocks[eid]
                user_block.append(etimestamp)
        for eid in range(numberOfUsers):
            offline_blocks[eid] = sorted(offline_blocks[eid])
        all_cnt = 0
        for event in events:
            ename, etimestamp_str, eids_str = event[0], event[1], event[2]
            if ename == "MESSAGE":
                etimestamp = int(etimestamp_str)
                if eids_str == "ALL":
                    all_cnt += 1
                elif eids_str == "HERE":
                    for eid in range(numberOfUsers):
                        user_offline_block = offline_blocks[eid]
                        is_offline = False
                        for offline_ts in user_offline_block:
                            if etimestamp >= offline_ts and etimestamp < (offline_ts + 60):
                                # user offline at this mention
                                is_offline = True
                                break
                        if not is_offline:
                            mentions[eid] += 1
                else:
                    eids_str_split = eids_str.split(" ")
                    for eid_str in eids_str_split:
                        eid = int(eid_str[2:])
                        mentions[eid] += 1
        for eid in range(numberOfUsers):
            mentions[eid] += all_cnt
        return mentions


def main() -> None:
    sol = Solution()
    result = sol.countMentions(
        2,
        # [["OFFLINE","10","0"],["MESSAGE","12","HERE"]]
        # [["MESSAGE","10","id1 id0"],["OFFLINE","11","0"],["MESSAGE","71","HERE"]]
        [["MESSAGE","10","id1 id0"],["OFFLINE","11","0"],["MESSAGE","12","ALL"]]
    )
    print(result)


main()
