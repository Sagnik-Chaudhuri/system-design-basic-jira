from classes.Subtrack import Subtrack
from classes.Task import Task


class Story(Task):

    allowed_status = ['open', 'in_progress', 'completed']

    def __init__(self, title, assignee, task_type, due_date, story_summary, status, subtracks=[]):
        Task.__init__(self, title=title, assignee=assignee,
                      task_type=task_type, due_date=due_date)
        self.story_summary = story_summary
        self.status = status if status in self.allowed_status else None
        self.subtracks = subtracks

    def change_task_assignee(self, new_assignee):
        return Task.change_task_assignee(self, new_assignee)

    def change_status(self, new_status):
        if new_status not in self.allowed_status:
            print('invalid story status')
            return
        if self.allowed_status.index(new_status) < self.allowed_status.index(self.status):
            print('new status cannot be behind prevous status')
            return

        if new_status != 'completed':
            self.status = new_status
        else:
            if len(self.subtracks) == 0:
                self.status = new_status
                print('\nchanging status of story from',
                      self.status, 'to ', new_status)
                return
            else:
                count_completed_subtracks = 0
                for subtrack in self.subtracks:
                    if subtrack.status == 'completed':
                        count_completed_subtracks += 1
                if len(self.subtracks) == count_completed_subtracks:
                    self.status = new_status
                    print('\nchanging status of story from',
                          self.status, 'to ', new_status)
                else:
                    print('cant complete story, subtracks still left')

        return

    def add_subtrack_to_current_story(self):
        if self.status == 'completed':
            print('\nstory is completed. cannot add subtrack')
        else:
            subtrack_title = input('subtrack_title?')
            subtrack_status = input('subtrack status?')
            subtrack = Subtrack(
                subtrack_title, status=subtrack_status, parent_task_title=self.title)
            self.subtracks.append(subtrack)

    def display_details(self):
        self.display_task()
        print('\nstory summary: ', self.story_summary, '\nstatus', self.status)
        if (len(self.subtracks) == 0):
            print('\ncurrent story has no sub_tracks')
        else:
            print('\ndisplaying subtracks for story')
            for subtrack in self.subtracks:
                subtrack.display_details()
