import prettytable

class create_table():
    def create_table(self, list_of_table_labels: list):
        table = prettytable.PrettyTable(list_of_table_labels)
        return table

    def add_rows_to_table(self, table_name, rows_of_table):
        try:
            if type(rows_of_table[0]) == list:
                for row in rows_of_table:
                    table_name.add_row(row)
                return table_name
            else:
                table_name.add_row(rows_of_table)
                return table_name
        except TypeError as e:
            print("TypeError: {0}".format(e))

    def create_entire_table(self, title:list, rows_of_table:list):
        lenth = len(title)
        for i in rows_of_table:
            if type(i) != list:
                if len(rows_of_table) != lenth:
                    raise ValueError("The length of rows you entered is not equal to the length of title.")
                else:
                    break
            else:
                if len(i) != lenth:
                    raise ValueError("The length of rows you entered is not equal to the length of title.")
        table = self.create_table(title)
        table = self.add_rows_to_table(table, rows_of_table)
        return table
    

print(create_table().create_entire_table([1,2],[3,4]))