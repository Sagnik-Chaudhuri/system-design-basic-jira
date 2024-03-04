from classes.Task import Task


class Feature(Task):

    allowed_status = ['open', 'in_progress', 'testing', 'deployed']

    def __init__(self, title, assignee, task_type, due_date, feature_summary, impact, status):
        Task.__init__(self, title, assignee, task_type, due_date)
        self.feature_summary = feature_summary
        self.impact = impact
        self.status = status if status in self.allowed_status else None

    def change_task_assignee(self, new_assignee):
        return Task.change_task_assignee(self, new_assignee)

    def change_status(self, new_status):
        if new_status not in self.allowed_status:
            print('invalid feature status')
            return
        if self.allowed_status.index(new_status) < self.allowed_status.index(self.status):
            print('new status cannot be behind prevous status')
            return

        print('\nchanging status of feature from',
              self.status, 'to ', new_status)
        self.status = new_status
        return

    def display_details(self):
        self.display_task()
        print('\nfeature summary: ', self.feature_summary,
              '\nimpact: ', self.impact, '\nstatus', self.status)
