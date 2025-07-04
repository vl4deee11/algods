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



// inclusion_exclusion:
// Считает мощность объединения множеств по обобщённому принципу включения-исключения.
//
// els — вектор объектов, каждый из которых определяет некоторое множество A_i (например, делители, фильтры, интервалы).
// and_f(subset) — функция, возвращающая мощность пересечения всех множеств, соответствующих элементам из subset:
//     and_f(S) = |A_{i1} ∩ A_{i2} ∩ ... ∩ A_{ik}|
//
// Возвращает: |A1 ∪ A2 ∪ ... ∪ An|
//
// ---
// Доказательство простыми словами:
//
// Если просто сложить мощности всех A_i, то элементы, попавшие в несколько A_i, будут посчитаны многократно.
//
// Чтобы учесть пересечения корректно, используется следующая формула:
//
// |A1 ∪ A2 ∪ ... ∪ An| =
//     ∑ |A_i|
//   - ∑ |A_i ∩ A_j|
//   + ∑ |A_i ∩ A_j ∩ A_k|
//   - ...
//   + (-1)^{k+1} * |A_{i1} ∩ A_{i2} ∩ ... ∩ A_{ik}|
//
// То есть:
// - Берём все непустые подмножества S ⊆ els,
// - Вычисляем пересечение множеств из этих элементов через and_f(S),
// - Включаем или исключаем в сумму по правилу: знак = (-1)^{|S| + 1}
// ---
template<typename T>
ll inclusion_exclusion(v(T)& els, function<long long(v(T)&)> and_f) {
    int n=els.size();
    ll tot=0;

    f(mask,1,1<<n){
        vector<T> sub;
        f(i,0,n){
            if (mask & (1 << i))
                sub.push_back(els[i]);
        }
        int sign=(__builtin_popcount(mask)%2==1)?1:-1;
        tot += sign*andF(sub);
    }
    return tot;
}
