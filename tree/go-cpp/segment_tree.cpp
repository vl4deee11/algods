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

using namespace std;
//#pragma GCC optimize("O3,unroll-loops")
//#pragma GCC target("avx2,bmi,bmi2,lzcnt,popcnt")

/* TYPES  */
#define ll long long
#define lli int64_t
#define str string
#define pii pair<int, int>
#define pll pair<long long, long long>
#define vi vector<int>
#define vs vector<string>
#define vll vector<long long>
#define mii map<int, int>
#define si set<int>
#define sc set<char>

/* FUNCTIONS */
#define v(t) vector<t>
#define s(t) set<t>
#define m(t, t2) map<t, t2>
#define um(t, t2) unordered_map<t, t2>
#define p(t, t2) pair<t, t2>
#define f(i, s, e) for(long long int i=s;i<e;i++)
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
ll min(ll a,int b) { if (a<b) return a; return b; }
ll min(int a,lli b) { if (a<b) return a; return b; }
ll min(ll a,ll b) { if (a<b) return a; return b; }
ll max(ll a,int b) { if (a>b) return a; return b; }
ll max(int a,ll b) { if (a>b) return a; return b; }
ll max(ll a,ll b) { if (a>b) return a; return b; }
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

/*  All Required define Pre-Processors and typedef Constants */
typedef long int int32;
typedef unsigned long int uint32;
typedef long long int int64;
typedef unsigned long long int uint64;

class ST {
public:
    ST(v(lli)& arr) {
        arrTree.resize(4 * arr.size());
        buildSegTree(arr, 1, 0, arr.size() - 1);
    }

    lli OpT(int l, int r) {
        return opT(1, 0, arrTree.size() / 4 - 1, l, r);
    }

    void Upd(int pos, lli newVal) {
        upd(1, 0, arrTree.size() / 4 - 1, pos, newVal);
    }

private:
    v(lli) arrTree;

    void buildSegTree(v(lli)& arr, int v, int tl, int tr) {
        if (tl == tr) {
            arrTree[v] = arr[tl];
            return;
        }
        int tm = (tl + tr) / 2;
        buildSegTree(arr, 2 * v, tl, tm);
        buildSegTree(arr, 2 * v + 1, tm + 1, tr);
        arrTree[v] = arrTree[2 * v] + arrTree[2 * v + 1];
    }

    lli opT(int v, int tl, int tr, int l, int r) {
        if (l > r) {
            return 0;
        }
        if (l == tl && r == tr) {
            return arrTree[v];
        }
        int tm = (tl + tr) / 2;
        return opT(2 * v, tl, tm, l, std::min(r, tm)) +
               opT(2 * v + 1, tm + 1, tr, std::max(l, tm + 1), r);
    }

    void upd(int v, int tl, int tr, int pos, lli newVal) {
        if (tl == tr) {
            arrTree[v] = newVal;
            return;
        }
        int tm = (tl + tr) / 2;
        if (pos <= tm) {
            upd(2 * v, tl, tm, pos, newVal);
        } else {
            upd(2 * v + 1, tm + 1, tr, pos, newVal);
        }
        arrTree[v] = arrTree[2 * v] + arrTree[2 * v + 1];
    }
};