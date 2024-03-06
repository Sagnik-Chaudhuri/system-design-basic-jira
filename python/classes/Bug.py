from classes.Task import Task


class Bug(Task):

    allowed_status = ['open', 'in_progress', 'fixed']

    def __init__(self, title, assignee, task_type, due_date, severity, status):
        Task.__init__(title, assignee, task_type, due_date)
        self.severity = severity
        self.status = status if status in self.allowed_status else None

    def change_task_assignee(self, new_assignee):
        return Task.change_task_assignee(self, new_assignee)

    def change_status(self, new_status):
        if new_status not in self.allowed_status:
            print('invalid bug status')
            return
        if self.allowed_status.index(new_status) < self.allowed_status.index(self.status):
            print('new status cannot be behind prevous status')
            return

        print('\nchanging status of bug from',
              self.status, 'to ', new_status)
        self.status = new_status
        return

    def display_details(self):
        self.display_task()
        print('\nbug severity: ', self.severity, '\nbug status', self.status)
