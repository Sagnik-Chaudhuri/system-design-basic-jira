class Task():
    def __init__(self, title, assignee, task_type, due_date):
        self.title = title
        self.assignee = assignee
        self.task_type = task_type
        self.due_date = due_date

    def change_task_assignee(self, new_assignee):
        self.assignee = new_assignee

    def display_task(self):
        print('\ntask title: ', self.title, '\ntask type: ',
              self.task_type, '\ndue date: ', self.due_date)
