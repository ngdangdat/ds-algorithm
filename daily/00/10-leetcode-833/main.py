from typing import Dict, List, Tuple

class Solution:
    def findReplaceString(self, s: str, indices: List[int], sources: List[str], targets: List[str]) -> str:
        # preprocess
        result = []
        source_target_map: Dict[int, List[Tuple[str, str]]] = {}
        for idx, indice in enumerate(indices):
            source_target_map.setdefault(indice, [])
            source_target_map[indice].append((sources[idx], targets[idx]))
        last_updated_index = 0
        for indice in sorted(source_target_map.keys()):
            source_target_list = source_target_map[indice]
            replace_text = None
            for source, target in source_target_list:
                k = len(source)
                if s[indice:indice+k] == source:
                    # proceed replace here
                    result.append(s[last_updated_index:indice])
                    last_updated_index = indice + len(source)
                    replace_text = target
                    break
            if replace_text is not None:
                result.append(replace_text)
        result.append(s[last_updated_index:len(s)])
        return "".join(result)


def main():
    # s = "abcd", indices[i] = 0, sources[i] = "ab", and targets[i] = "eee"
    sol = Solution()
    testcases = [
        ("abcd", [0], ["ab"], ["eee"], "eeecd"),
        ("abcd", [0, 2], ["a", "cd"], ["eee", "ffff"], "eeebffff"),
        ("abcd", [0, 2], ["ab", "ec"], ["eee", "ffff"], "eeecd"),
        ("vmokgggqzp", [3, 5, 1], ["kg","ggq","mo"], ["s","so","bfr"], "vbfrssozp"),
        ("abcd", [0, 0], ["a", "b"], ["b", "c"], "bbcd"),
        ("abcde", [2, 2], ["bc", "cde"], ["fe", "f"], "abf"),
        ("jjievdtjfb", [4, 6, 1], ["md","tjgb","jf"], ["foe","oov","e"], "jjievdtjfb"),
    ]

    for idx, tc in enumerate(testcases):
        print(f"CASE {idx+1}")
        s = tc[0]
        indicies = tc[1]
        sources = tc[2]
        targets = tc[3]
        expected = tc[4]
        result = sol.findReplaceString(s, indicies, sources, targets)
        res = "FAILED"
        if result == expected:
            res = "PASSED"
        print(f"{res} expect={expected}, got={result}")


if __name__ == "__main__":
    main()
