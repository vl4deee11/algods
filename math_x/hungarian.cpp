#include <iostream>
#include <queue>
#include <vector>
#include <cmath>
#include <string>
#include <sstream>
#include <cstring>
#include <algorithm>
#include <climits>
#include <stack>
#include <set>
#include <map>
#include <list>
#include <cassert>
#include <unordered_map>
#include <numeric>
#include <iomanip>
#include <iostream>
#include <cstdio>
#include <algorithm>
#include <queue>
#include <cstdlib>
#include <cstring>
#include <bitset>

using namespace std;
#pragma GCC optimize("O3,unroll-loops")
#pragma GCC target("avx2,bmi,bmi2,lzcnt,popcnt")

/* TYPES  */
#define xf first
#define xs second
#define ll long long int
#define ul unsigned long long int
#define lli int64_t
#define ulli uint64_t
#define dbl double
#define ldbl long double
#define str string
#define pii pair<int, int>
#define pll pair<ll,ll>
#define vi vector<int>
#define vs vector<string>
#define vll vector<long long>
#define mii map<int, int>
#define si set<int>
#define sc set<char>

/* FUNCTIONS */
#define mid(l, r) 	       ((l + r) >> 1)
#define all(a)             a.begin(),a.end()
#define v(t) vector<t>
#define st(t) stack<t>
#define ar(t,sz) array<t,sz>
#define s(t) set<t>
#define ss(a) sort(a.begin(),a.end())
#define ms(t) multiset<t>
#define mipq(t) priority_queue<t,v(t),greater<t>>
#define mapq(t) priority_queue<t,v(t),less<t>>
#define trpl(a,b,c) tuple<a,b,c>
#define m(t, t2) map<t, t2>
#define um(t, t2) unordered_map<t, t2>
#define p(t, t2) pair<t, t2>
#define f(i, s, e) for(long long int i=s;i<e;i++)
#define fc(i, s, e, c) for(long long int i=s;i<e;i+=c)
#define fa(k,in) for(auto k:in)
#define fm(i, s, e) for(long long int i=s;i!=e;i++)
#define cf(i, s, e) for(long long int i=s;i<=e;i++)
#define rf(i, e, s) for(long long int i=e-1;i>=s;i--)
#define wl(c) while(c)
#define pb push_back
#define mpp make_pair
#define cl clear
#define eb emplace_back

/* PRINTS */
template <class T>
void print_v(vector<T> &v) { cout << "{"; for (auto x : v) cout << x << ","; cout << "\b}"; }

/* UTILS */
#define MOD 1000000007
#define PI 3.1415926535897932384626433832795
#define read(type) readInt<type>()
//ll min(ll a,int b) { if (a<b) return a; return b; }
//ll min(int a,ll b) { if (a<b) return a; return b; }
lli min(lli a,lli b) { if (a<b) return a; return b; }
//ll max(ll a,int b) { if (a>b) return a; return b; }
//ll max(int a,ll b) { if (a>b) return a; return b; }
lli max(lli a,lli b) { if (a>b) return a; return b; }
int chaz_to_int026(char x) {return int(x - 'a');}
int chAZ_to_int026(char x) {return int(x - 'A');}
char int026_to_chaz(int x) {return char(x + 'a');}
char int026_to_chAZ(int x) {return char(x + 'A');}
int mod(int a, int b) {return (a % b + b) % b;}
bool float64_eq(double f1, double f2) {return abs(f1 - f2) < 1e-6;}
bool float64_gt_or_eq(double f1, double f2) {return float64_eq(f1, f2) || f1 > f2;}
bool float64_lt_or_eq(double f1, double f2) {return float64_eq(f1, f2) || f1 < f2;}
ll sum_ap(ll a, ll b) {ll n = (b - a) + 1;return (n * (a + b)) / 2;}
double dist(pii p1, pii p2) {return sqrt(pow(double(p2.first - p1.first), 2) + pow(double(p2.second - p1.second), 2));}
ll gcd(ll a,ll b) { if (b==0) return a; return gcd(b, a%b); }
ll lcm(ll a,ll b) { return a/gcd(a,b)*b; }
pair<int, int> chess(string b) {return make_pair(7 - (int(b[0]) - 48) - 1, 7 - (int(b[1]) - 'a'));}
bool ch(pair<int, int> a, pair<int, int> b, pair<int, int> c) {return int64_t(b.first - a.first) * int64_t(c.second - a.second) - int64_t(b.second - a.second) * int64_t(c.first - a.first) >= 0;}
string to_upper(string a) { for (int i=0;i<(int)a.size();++i) if (a[i]>='a' && a[i]<='z') a[i]-='a'-'A'; return a; }
string to_lower(string a) { for (int i=0;i<(int)a.size();++i) if (a[i]>='A' && a[i]<='Z') a[i]+='a'-'A'; return a; }
bool prime(ll a) { if (a==1) return 0; for (int i=2;i<=round(sqrt(a));++i) if (a%i==0) return 0; return 1; }
vector<string> split(const string& s, char delimiter){vector<string> tokens;string token;istringstream tokenStream(s);wl(std::getline(tokenStream, token, delimiter)){tokens.push_back(token);};return tokens;}
ll mod_exp(ll b,ll e,ll mod){ll r=1;b=b%mod;wl(e>0){if(e%2==1){r=(r*b)%mod;};e/=2;b=(b*b)%mod;}return r;}


template <typename T>
struct hungarian {  // km
    int n;
    vector<int> matchx;
    vector<int> matchy;
    vector<int> pre;
    vector<bool> visx;
    vector<bool> visy;
    vector<T> lx;
    vector<T> ly;
    vector<vector<T> > g;
    vector<T> slack;
    T inf;
    T res;
    queue<int> q;
    int org_n;
    int org_m;

    hungarian(int _n, int _m) {
        org_n = _n;
        org_m = _m;
        n = max(_n, _m);
        inf = numeric_limits<T>::max();
        res = 0;
        g = vector<vector<T> >(n, vector<T>(n));
        matchx = vector<int>(n, -1);
        matchy = vector<int>(n, -1);
        pre = vector<int>(n);
        visx = vector<bool>(n);
        visy = vector<bool>(n);
        lx = vector<T>(n, -inf);
        ly = vector<T>(n);
        slack = vector<T>(n);
    }

    void addEdge(int u, int v, ll w) {
        g[u][v] = max(w, 0);
    }

    bool check(int v) {
        visy[v] = true;
        if (matchy[v] != -1) {
            q.push(matchy[v]);
            visx[matchy[v]] = true;  // in S
            return false;
        }

        while (v != -1) {
            matchy[v] = pre[v];
            swap(v, matchx[pre[v]]);
        }
        return true;
    }

    void bfs(int i) {
        while (!q.empty()) {
            q.pop();
        }
        q.push(i);
        visx[i] = true;
        while (true) {
            while (!q.empty()) {
                int u = q.front();
                q.pop();
                for (int v = 0; v < n; v++) {
                    if (!visy[v]) {
                        T delta = lx[u] + ly[v] - g[u][v];
                        if (slack[v] >= delta) {
                            pre[v] = u;
                            if (delta) {
                                slack[v] = delta;
                            } else if (check(v)) {
                                return;
                            }
                        }
                    }
                }
            }

            T a = inf;
            for (int j = 0; j < n; j++) {
                if (!visy[j]) {
                    a = min(a, slack[j]);
                }
            }
            for (int j = 0; j < n; j++) {
                if (visx[j]) {  // S
                    lx[j] -= a;
                }
                if (visy[j]) {  // T
                    ly[j] += a;
                } else {  // T'
                    slack[j] -= a;
                }
            }
            for (int j = 0; j < n; j++) {
                if (!visy[j] && slack[j] == 0 && check(j)) {
                    return;
                }
            }
        }
    }

    int solve() {
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                lx[i] = max(lx[i], g[i][j]);
            }
        }

        for (int i = 0; i < n; i++) {
            fill(slack.begin(), slack.end(), inf);
            fill(visx.begin(), visx.end(), false);
            fill(visy.begin(), visy.end(), false);
            bfs(i);
        }

        return res;
    }
};