import unittest
from typing import Union

# https://cp-algorithms.com/data_structures/treap.html#merge


class TreapNode:
    def __init__(self, key: int = None, prior: int = None):
        self.l: Union[TreapNode, None] = None
        self.r: Union[TreapNode, None] = None
        self.key = key
        self.prior = prior

    def __repr__(self):
        return f"{self.key} - {self.prior}"


def insert(root: TreapNode, node: TreapNode) -> TreapNode:
    """
    :param root: корень
    :param node: новый узел
    :return: новый корень
    """
    # Вставка в пустой узел. Просто вернем вставляемый (рекурсивный выход)
    if root is None:
        return node
    # Если приоритет узла больше чем у корня, то новым корня станет этот вставляемый узел.
    # Поэтому нужно "взвесить" все за него
    # Поэтому возьмем ключ узла и пихнем в его левое и правое дерево данные из корня.
    # В левой ветке окажутся узлы с <= key, в правой с большими
    elif node.prior > root.prior:
        node.l, node.r = split(root, node.l, node.r, node.key)
        return node
    else:
        # Если с приоритетом все норм, тупо решаем по правило построения бинарного дерева куда пихать новый узел
        if root.key <= node.key:
            root.r = insert(root.r, node)
        else:
            root.l = insert(root.l, node)
        return root


def unite(l: TreapNode, r: TreapNode):
    # Если один из узлов None, вернем другой
    if l is None or r is None:
        if l is None:
            return r
        return l
    # За новый корень дерева будем считать левое (первое поддерево)
    # Проверим все ли ок с приоритетами для соблюдения свойства кучи, если что свапнем
    if l.prior < r.prior:
        l, r = r, l
    # создадим пустые ноды, чтобы наполнить их узлами с помощью сплит функции
    l_root, r_root = TreapNode(), TreapNode()
    # Сделаем сплит по всему r дереву, используя l.key ключ.
    # Получим в l_root все что меньше l.key, в r_root все что больше
    l_root, r_root = split(r, l_root, r_root, l.key)
    # Провернем то же самое для левого и правого ребенка узла l, используя найденные левые и правые поддеревья из r
    l.l = unite(l.l, l_root)
    l.r = unite(l.r, r_root)
    return l


def split(root: TreapNode, l: TreapNode, r: TreapNode, key: int):
    if root is None:
        # Вернем пустые l и r
        return None, None
    # Если ключ корня меньше либо равен тому, по которому сплитуем, то...
    elif root.key <= key:
        # Нам нужно чтобы в левом корне было все из root.l, а также все корни <= key из правой ветки.
        # Делаем так потому что все они гарантированно меньше key
        # в правом корне будут все остальные ключи из правой ветки (гарантированно больше key)
        root.r, r = split(root.r, root.r, r, key)
        return root, r
    else:
        # Здесь ровно наоборот....
        # Т.к. key меньше чем ключ root, то слева должны оказаться все узлы с меньшим key
        # в правом корне окажется все из root.r, а также те узлы, N.key которых больше чем key
        l, root.l, = split(root.l, l, root.l, key)
        return l, root


class TestTreap(unittest.TestCase):
    def test_treap_building(self):
        n1 = TreapNode(2, 2)
        n2 = TreapNode(5, 6)
        n3 = TreapNode(10, 1)
        n4 = TreapNode(18, 1)
        n5 = TreapNode(19, 2)
        n6 = TreapNode(2, 9)
        n7 = TreapNode(18, 11)
        n8 = TreapNode(1, 14)
        n9 = TreapNode(14, 14)
        n10 = TreapNode(6, 15)

        root = unite(n1, n2)
        root = insert(root, n3)
        root = insert(root, n4)
        root = insert(root, n5)
        root = insert(root, n6)
        root = insert(root, n7)
        root = insert(root, n8)
        root = insert(root, n9)
        root = insert(root, n10)

        self.assertTrue(root.prior == 15)
        self.assertTrue(root.l.prior == 14)
        self.assertTrue(root.l.r.prior == 9)
        self.assertTrue(root.l.r.l.prior == 2)
        self.assertTrue(root.l.r.r.prior == 6)

        self.assertTrue(root.r.prior == 14)
        self.assertTrue(root.r.l.prior == 1)
        self.assertTrue(root.r.r.prior == 11)
        self.assertTrue(root.r.r.l.prior == 1)
        self.assertTrue(root.r.r.r.prior == 2)



