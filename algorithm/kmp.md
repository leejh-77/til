KMP 는 문자열에서 특정 문자열을 찾아내는 알고리즘이다.
핵심은 문자열 비교를 하다가 틀린 문자를 발견했을 때, 얼마나 앞으로 다시 이동해야하는지에 대한 테이블을 따로 갖고있는 것이다.  

```java
private static int[] kmp_table(String p) {
    int m = p.length();
    int begin = 1, matched = 0;
    int[] pi = new int[m];

    while (begin + matched < m) {
        if (p.charAt(begin + matched) == p.charAt(matched)) {
            matched++;
            pi[begin+matched-1] = matched;
        } else {
            if (matched == 0) {
                begin++;
            } else {
                begin += matched - pi[matched-1];
                matched = pi[matched-1];
            }
        }
    }

    return pi;
}

private static int kmp(String s, String p) {
    int n = s.length(), m = p.length();
    int[] pi = kmp_table(p);
    int begin = 0, matched = 0;

    while (begin <= n - m) {
        if (matched < m && s.charAt(begin + matched) == p.charAt(matched)) {
            matched++;
            if (matched == m) {
                return begin+matched;
            }
        } else {
            if (matched == 0) {
                begin++;
            } else {
                begin += matched - pi[matched-1];
                matched = pi[matched-1];
            }
        }
    }
    return -1;
}
```